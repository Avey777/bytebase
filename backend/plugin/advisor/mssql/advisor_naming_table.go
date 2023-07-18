// Package mssql is the advisor for MSSQL database.
package mssql

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tsql-parser"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common/log"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
	bbparser "github.com/bytebase/bytebase/backend/plugin/parser/sql"
)

var (
	_ advisor.Advisor = (*NamingTableAdvisor)(nil)
)

func init() {
	advisor.Register(db.MSSQL, advisor.MSSQLNamingTableConvention, &NamingTableAdvisor{})
}

// NamingTableAdvisor is the advisor checking for table naming convention..
type NamingTableAdvisor struct {
}

// Check checks for table naming convention..
func (*NamingTableAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	tree, ok := ctx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	format, maxLength, err := advisor.UnmarshalNamingRulePayloadAsRegexp(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}

	listener := &namingTableListener{
		level:     level,
		title:     string(ctx.Rule.Type),
		format:    format,
		maxLength: maxLength,
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// namingTableListener is the listener for table naming convention.
type namingTableListener struct {
	*parser.BaseTSqlParserListener

	level  advisor.Status
	title  string
	format *regexp.Regexp
	// maxLength is the max length of the table name.
	maxLength int

	adviceList []advisor.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *namingTableListener) generateAdvice() ([]advisor.Advice, error) {
	if len(l.adviceList) == 0 {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return l.adviceList, nil
}

// EnterCreate_table is called when production create_table is entered.
func (l *namingTableListener) EnterCreate_table(ctx *parser.Create_tableContext) {
	tableName := ctx.Table_name().GetTable().GetText()

	if !l.format.MatchString(tableName) {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NamingTableConventionMismatch,
			Title:   l.title,
			Content: fmt.Sprintf(`%s mismatches table naming convention, naming format should be %q`, tableName, l.format),
			Line:    ctx.GetStart().GetLine(),
		})
	}
	if l.maxLength > 0 && len(tableName) > l.maxLength {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NamingTableConventionMismatch,
			Title:   l.title,
			Content: fmt.Sprintf(`%s mismatches table naming convention, its length should be within %d characters`, tableName, l.maxLength),
			Line:    ctx.GetStart().GetLine(),
		})
	}
}

// EnterExecute_body is called when production execute_body is entered.
func (l *namingTableListener) EnterExecute_body(ctx *parser.Execute_bodyContext) {
	if ctx.Func_proc_name_server_database_schema() == nil {
		return
	}
	if ctx.Func_proc_name_server_database_schema().Func_proc_name_database_schema() == nil {
		return
	}
	if ctx.Func_proc_name_server_database_schema().Func_proc_name_database_schema().Func_proc_name_schema() == nil {
		return
	}
	if ctx.Func_proc_name_server_database_schema().Func_proc_name_database_schema().Func_proc_name_schema().GetSchema() != nil {
		return
	}

	v := ctx.Func_proc_name_server_database_schema().Func_proc_name_database_schema().Func_proc_name_schema().GetProcedure()
	normalizedProcedureName, err := bbparser.NormalizedTSqlTableNamePart(v)
	if err != nil {
		log.Error(errors.Wrapf(err, "failed to normalize procedure name").Error())
		return
	}
	if normalizedProcedureName != "sp_rename" {
		return
	}

	firstArgument := ctx.Execute_statement_arg()
	if firstArgument == nil {
		return
	}
	if firstArgument.Execute_statement_arg_unnamed() == nil {
		return
	}
	if firstArgument.Execute_statement_arg_unnamed().Execute_parameter() == nil {
		return
	}
	if firstArgument.Execute_statement_arg_unnamed().Execute_parameter().Constant() == nil {
		return
	}
	if firstArgument.Execute_statement_arg_unnamed().Execute_parameter().Constant().STRING() == nil {
		return
	}

	if len(ctx.Execute_statement_arg().AllExecute_statement_arg()) != 1 {
		return
	}
	secondArgument := ctx.Execute_statement_arg().Execute_statement_arg(0)
	if secondArgument == nil {
		return
	}
	if secondArgument.Execute_statement_arg_unnamed() == nil {
		return
	}
	if secondArgument.Execute_statement_arg_unnamed().Execute_parameter() == nil {
		return
	}
	if secondArgument.Execute_statement_arg_unnamed().Execute_parameter().Constant() == nil {
		return
	}
	if secondArgument.Execute_statement_arg_unnamed().Execute_parameter().Constant().STRING() == nil {
		return
	}

	newTableName := secondArgument.Execute_statement_arg_unnamed().Execute_parameter().Constant().STRING().GetText()
	if strings.HasPrefix(newTableName, "'") && strings.HasSuffix(newTableName, "'") {
		newTableName = newTableName[1 : len(newTableName)-1]
	}

	if !l.format.MatchString(newTableName) {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NamingTableConventionMismatch,
			Title:   l.title,
			Content: fmt.Sprintf(`%s mismatches table naming convention, naming format should be %q`, newTableName, l.format),
			Line:    ctx.GetStart().GetLine(),
		})
	}
	if l.maxLength > 0 && len(newTableName) > l.maxLength {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NamingTableConventionMismatch,
			Title:   l.title,
			Content: fmt.Sprintf(`%s mismatches table naming convention, its length should be within %d characters`, newTableName, l.maxLength),
			Line:    ctx.GetStart().GetLine(),
		})
	}
}

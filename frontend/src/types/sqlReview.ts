import { t } from "@/plugins/i18n";
import type { Engine } from "@/types/proto/v1/common";
import { engineFromJSON } from "@/types/proto/v1/common";
import {
  SQLReviewRuleLevel,
  sQLReviewRuleLevelFromJSON,
} from "@/types/proto/v1/org_policy_service";
import type { PlanType } from "@/types/proto/v1/subscription_service";
import sqlReviewSchema from "./sql-review-schema.yaml";
import sqlReviewDevTemplate from "./sql-review.dev.yaml";
import sqlReviewProdTemplate from "./sql-review.prod.yaml";
import sqlReviewSampleTemplate from "./sql-review.sample.yaml";

export const LEVEL_LIST = [
  SQLReviewRuleLevel.ERROR,
  SQLReviewRuleLevel.WARNING,
  SQLReviewRuleLevel.DISABLED,
];

// NumberPayload is the number type payload configuration options and default value.
// Used by the frontend.
export interface NumberPayload {
  type: "NUMBER";
  default: number;
  value?: number;
}

// StringPayload is the string type payload configuration options and default value.
// Used by the frontend.
export interface StringPayload {
  type: "STRING";
  default: string;
  value?: string;
}

// BooleanPayload is the boolean type payload configuration options and default value.
// Used by the frontend.
export interface BooleanPayload {
  type: "BOOLEAN";
  default: boolean;
  value?: boolean;
}

// StringArrayPayload is the string array type payload configuration options and default value.
// Used by the frontend.
export interface StringArrayPayload {
  type: "STRING_ARRAY";
  default: string[];
  value?: string[];
}

// TemplatePayload is the string template type payload configuration options and default value.
// Used by the frontend.
export interface TemplatePayload {
  type: "TEMPLATE";
  default: string;
  templateList: string[];
  value?: string;
}

// RuleConfigComponent is the rule configuration options and default value.
// Used by the frontend.
export interface RuleConfigComponent {
  key: string;
  payload:
    | StringPayload
    | NumberPayload
    | TemplatePayload
    | StringArrayPayload
    | BooleanPayload;
}

// The naming format rule payload.
// Used by the backend.
interface NamingFormatPayload {
  format: string;
  maxLength?: number;
}

// The string array rule payload.
// Used by the backend.
interface StringArrayLimitPayload {
  list: string[];
}

// The comment format rule payload.
// Used by the backend.
interface CommentFormatPayload {
  required: boolean;
  maxLength: number;
}

// The number value rule payload.
// Used by the backend.
interface NumberValuePayload {
  number: number;
}

// The string value rule payload.
// Used by the backend.
interface StringValuePayload {
  string: string;
}

// The case rule payload.
// Used by the backend.
interface CasePayload {
  upper: boolean;
}

// The SchemaPolicyRule stores the rule configuration by users.
// Used by the backend
export interface SchemaPolicyRule {
  type: string;
  level: SQLReviewRuleLevel;
  engine: Engine;
  payload?:
    | NamingFormatPayload
    | StringArrayLimitPayload
    | CommentFormatPayload
    | NumberValuePayload
    | StringValuePayload
    | CasePayload;
  comment: string;
}

// The API for SQL review policy in backend.
export interface SQLReviewPolicy {
  id: string;

  // Standard fields
  // enforce means if the policy is active
  enforce: boolean;

  // Domain specific fields
  name: string;
  ruleList: SchemaPolicyRule[];
  resources: string[];
}

// RuleTemplateV2 is the rule template. Used by the frontend
export interface RuleTemplateV2 {
  type: string;
  category: string;
  engine: Engine;
  level: SQLReviewRuleLevel;
  componentList: RuleConfigComponent[];
  comment?: string;
}

// SQLReviewPolicyTemplateV2 is the rule template set
export interface SQLReviewPolicyTemplateV2 {
  id: string;
  ruleList: RuleTemplateV2[];
}

export const ruleTemplateMapV2 = (sqlReviewSchema as RuleTemplateV2[]).reduce(
  (map, ruleSchema) => {
    if (!map.has(ruleSchema.type)) {
      map.set(ruleSchema.type, new Map());
    }
    const engine = engineFromJSON(ruleSchema.engine);
    map.get(ruleSchema.type)?.set(engine, {
      ...ruleSchema,
      engine,
      componentList: ruleSchema.componentList || [],
    });
    return map;
  },
  new Map<string, Map<Engine, RuleTemplateV2>>()
);

// Build the frontend template list based on schema and template.
export const TEMPLATE_LIST_V2: SQLReviewPolicyTemplateV2[] = (function () {
  interface PayloadObject {
    [key: string]: any;
  }
  const templateList = [
    sqlReviewSampleTemplate,
    sqlReviewDevTemplate,
    sqlReviewProdTemplate,
  ] as {
    id: string;
    ruleList: {
      type: string;
      level: SQLReviewRuleLevel;
      payload?: PayloadObject;
      engine: Engine;
    }[];
  }[];

  return templateList.map((template) => {
    const ruleList: RuleTemplateV2[] = [];

    for (const rule of template.ruleList) {
      const ruleTemplate = ruleTemplateMapV2
        .get(rule.type)
        ?.get(engineFromJSON(rule.engine));
      if (!ruleTemplate) {
        continue;
      }

      ruleList.push({
        ...ruleTemplate,
        level: sQLReviewRuleLevelFromJSON(rule.level),
        // Using template rule payload to override the component list.
        componentList: ruleTemplate.componentList.map((component) => {
          if (rule.payload && rule.payload[component.key]) {
            return {
              ...component,
              payload: {
                ...component.payload,
                default: rule.payload[component.key],
              },
            };
          }
          return component;
        }),
      });
    }

    return {
      id: template.id,
      ruleList,
    };
  });
})();

interface RuleCategory {
  id: string;
  ruleList: RuleTemplateV2[];
}

// convertToCategoryList will reduce RuleTemplate list to RuleCategory list.
export const convertToCategoryList = (
  ruleList: RuleTemplateV2[]
): RuleCategory[] => {
  const dict = ruleList.reduce(
    (dict, rule) => {
      if (!dict[rule.category]) {
        dict[rule.category] = {
          id: rule.category,
          ruleList: [],
        };
      }
      dict[rule.category].ruleList.push(rule);
      return dict;
    },
    {} as { [key: string]: RuleCategory }
  );

  return Object.values(dict);
};

// The convertPolicyRuleToRuleTemplate will convert the review policy rule to rule template for frontend useage.
// Will throw exception if we don't implement the payload handler for specific type of rule.
export const convertPolicyRuleToRuleTemplate = (
  policyRule: SchemaPolicyRule,
  ruleTemplate: RuleTemplateV2
): RuleTemplateV2 => {
  if (policyRule.type !== ruleTemplate.type) {
    throw new Error(
      `The rule type is not same. policyRule:${ruleTemplate.type}, ruleTemplate:${ruleTemplate.type}`
    );
  }

  const res: RuleTemplateV2 = {
    ...ruleTemplate,
    engine: policyRule.engine,
    level: policyRule.level,
    comment: policyRule.comment,
  };

  const componentList = ruleTemplate.componentList ?? [];
  if (componentList.length === 0) {
    return res;
  }

  const payload = policyRule?.payload ?? {};

  const stringComponent = componentList.find(
    (c) => c.payload.type === "STRING"
  );
  const numberComponent = componentList.find(
    (c) => c.payload.type === "NUMBER"
  );
  const booleanComponent = componentList.find(
    (c) => c.payload.type === "BOOLEAN"
  );
  const templateComponent = componentList.find(
    (c) => c.payload.type === "TEMPLATE"
  );
  const stringArrayComponent = componentList.find(
    (c) => c.payload.type === "STRING_ARRAY"
  );

  switch (ruleTemplate.type) {
    case "statement.query.minimum-plan-level":
      if (!stringComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }

      return {
        ...res,
        componentList: [
          {
            ...stringComponent,
            payload: {
              ...stringComponent.payload,
              value: (payload as StringValuePayload).string,
            } as StringPayload,
          },
        ],
      };
    // Following rules require STRING component.
    case "table.drop-naming-convention":
      if (!stringComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }

      return {
        ...res,
        componentList: [
          {
            ...stringComponent,
            payload: {
              ...stringComponent.payload,
              value: (payload as NamingFormatPayload).format,
            } as StringPayload,
          },
        ],
      };
    // Following rules require STRING and NUMBER component.
    case "naming.column":
    case "naming.column.auto-increment":
    case "naming.table":
      if (!stringComponent || !numberComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }

      return {
        ...res,
        componentList: [
          {
            ...stringComponent,
            payload: {
              ...stringComponent.payload,
              value: (payload as NamingFormatPayload).format,
            } as StringPayload,
          },
          {
            ...numberComponent,
            payload: {
              ...numberComponent.payload,
              value: (payload as NamingFormatPayload).maxLength,
            } as NumberPayload,
          },
        ],
      };
    // Following rules require TEMPLATE component.
    case "naming.index.pk":
      if (!templateComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }

      return {
        ...res,
        componentList: [
          {
            ...templateComponent,
            payload: {
              ...templateComponent.payload,
              value: (payload as NamingFormatPayload).format,
            } as TemplatePayload,
          },
        ],
      };
    // Following rules require TEMPLATE and NUMBER component.
    case "naming.index.idx":
    case "naming.index.uk":
    case "naming.index.fk":
      if (!templateComponent || !numberComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }

      return {
        ...res,
        componentList: [
          {
            ...templateComponent,
            payload: {
              ...templateComponent.payload,
              value: (payload as NamingFormatPayload).format,
            } as TemplatePayload,
          },
          {
            ...numberComponent,
            payload: {
              ...numberComponent.payload,
              value: (payload as NamingFormatPayload).maxLength,
            } as NumberPayload,
          },
        ],
      };
    // Following rules require BOOLEAN component.
    case "naming.identifier.case":
      if (!booleanComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }
      return {
        ...res,
        componentList: [
          {
            ...booleanComponent,
            payload: booleanComponent.payload,
          },
        ],
      };
    case "column.required": {
      const requiredColumnComponent = componentList[0];
      // The columnList payload is deprecated.
      // Just keep it to compatible with old data, we can remove this later.
      let value: string[] = (payload as any)["columnList"];
      if (!value) {
        value = (payload as StringArrayLimitPayload).list;
      }
      const requiredColumnPayload = {
        ...requiredColumnComponent.payload,
        value,
      } as StringArrayPayload;
      return {
        ...res,
        componentList: [
          {
            ...requiredColumnComponent,
            payload: requiredColumnPayload,
          },
        ],
      };
    }
    // Following rules require STRING_ARRAY component.
    case "column.type-disallow-list":
    case "index.primary-key-type-allowlist":
    case "index.type-allow-list":
    case "system.charset.allowlist":
    case "system.collation.allowlist":
    case "system.function.disallowed-list":
    case "table.disallow-dml":
    case "table.disallow-ddl": {
      if (!stringArrayComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }
      const stringArrayPayload = {
        ...stringArrayComponent.payload,
        value: (payload as StringArrayLimitPayload).list,
      } as StringArrayPayload;
      return {
        ...res,
        componentList: [
          {
            ...stringArrayComponent,
            payload: stringArrayPayload,
          },
        ],
      };
    }
    // Following rules require BOOLEAN and NUMBER component.
    case "column.comment":
    case "table.comment":
      if (!booleanComponent || !numberComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }

      return {
        ...res,
        componentList: [
          {
            ...booleanComponent,
            payload: {
              ...booleanComponent.payload,
              value: (payload as CommentFormatPayload).required,
            } as BooleanPayload,
          },
          {
            ...numberComponent,
            payload: {
              ...numberComponent.payload,
              value: (payload as CommentFormatPayload).maxLength,
            } as NumberPayload,
          },
        ],
      };
    // Following rules require NUMBER component.
    case "statement.insert.row-limit":
    case "statement.affected-row-limit":
    case "column.maximum-character-length":
    case "column.maximum-varchar-length":
    case "column.auto-increment-initial-value":
    case "index.key-number-limit":
    case "index.total-number-limit":
    case "system.comment.length":
    case "advice.online-migration":
    case "table.text-fields-total-length":
    case "table.limit-size":
    case "statement.where.maximum-logical-operator-count":
    case "statement.maximum-limit-value":
    case "statement.maximum-join-table-count":
    case "statement.maximum-statements-in-transaction":
      if (!numberComponent) {
        throw new Error(`Invalid rule ${ruleTemplate.type}`);
      }

      return {
        ...res,
        componentList: [
          {
            ...numberComponent,
            payload: {
              ...numberComponent.payload,
              value: (payload as NumberValuePayload).number,
            } as NumberPayload,
          },
        ],
      };
  }

  throw new Error(`Invalid rule ${ruleTemplate.type}`);
};

const mergeIndividualConfigAsRule = (
  base: SchemaPolicyRule,
  template: RuleTemplateV2
): SchemaPolicyRule => {
  const componentList = template.componentList ?? [];
  const stringPayload = componentList.find((c) => c.payload.type === "STRING")
    ?.payload as StringPayload | undefined;
  const numberPayload = componentList.find((c) => c.payload.type === "NUMBER")
    ?.payload as NumberPayload | undefined;
  const templatePayload = componentList.find(
    (c) => c.payload.type === "TEMPLATE"
  )?.payload as TemplatePayload | undefined;
  const booleanPayload = componentList.find((c) => c.payload.type === "BOOLEAN")
    ?.payload as BooleanPayload | undefined;
  const stringArrayPayload = componentList.find(
    (c) => c.payload.type === "STRING_ARRAY"
  )?.payload as StringArrayPayload | undefined;

  switch (template.type) {
    case "statement.query.minimum-plan-level":
      if (!stringPayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }

      return {
        ...base,
        payload: {
          string: stringPayload.value ?? stringPayload.default,
        },
      };
    // Following rules require STRING component.
    case "table.drop-naming-convention":
      if (!stringPayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }

      return {
        ...base,
        payload: {
          format: stringPayload.value ?? stringPayload.default,
        },
      };
    // Following rules require STRING and NUMBER component.
    case "naming.column":
    case "naming.column.auto-increment":
    case "naming.table":
      if (!stringPayload || !numberPayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }

      return {
        ...base,
        payload: {
          format: stringPayload.value ?? stringPayload.default,
          maxLength: numberPayload.value ?? numberPayload.default,
        },
      };
    // Following rules require TEMPLATE component.
    case "naming.index.pk":
      if (!templatePayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }
      return {
        ...base,
        payload: {
          format: templatePayload.value ?? templatePayload.default,
        },
      };
    // Following rules require TEMPLATE and NUMBER component.
    case "naming.index.idx":
    case "naming.index.uk":
    case "naming.index.fk":
      if (!templatePayload || !numberPayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }
      return {
        ...base,
        payload: {
          format: templatePayload.value ?? templatePayload.default,
          maxLength: numberPayload.value ?? numberPayload.default,
        },
      };
    // Following rules require BOOLEAN component.
    case "naming.identifier.case":
      if (!booleanPayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }
      return {
        ...base,
        payload: {
          upper: booleanPayload.value ?? booleanPayload.default,
        },
      };
    // Following rules require STRING_ARRAY component.
    case "column.required":
    case "column.type-disallow-list":
    case "index.primary-key-type-allowlist":
    case "index.type-allow-list":
    case "system.charset.allowlist":
    case "system.collation.allowlist":
    case "system.function.disallowed-list":
    case "table.disallow-dml":
    case "table.disallow-ddl": {
      if (!stringArrayPayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }
      return {
        ...base,
        payload: {
          list: stringArrayPayload.value ?? stringArrayPayload.default,
        },
      };
    }
    // Following rules require BOOLEAN and NUMBER component.
    case "column.comment":
    case "table.comment":
      if (!booleanPayload || !numberPayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }
      return {
        ...base,
        payload: {
          required: booleanPayload.value ?? booleanPayload.default,
          maxLength: numberPayload.value ?? numberPayload.default,
        },
      };
    // Following rules require NUMBER component.
    case "statement.insert.row-limit":
    case "statement.affected-row-limit":
    case "column.maximum-character-length":
    case "column.maximum-varchar-length":
    case "column.auto-increment-initial-value":
    case "index.key-number-limit":
    case "index.total-number-limit":
    case "system.comment.length":
    case "advice.online-migration":
    case "table.text-fields-total-length":
    case "table.limit-size":
    case "statement.where.maximum-logical-operator-count":
    case "statement.maximum-limit-value":
    case "statement.maximum-join-table-count":
    case "statement.maximum-statements-in-transaction":
      if (!numberPayload) {
        throw new Error(`Invalid rule ${template.type}`);
      }
      return {
        ...base,
        payload: {
          number: numberPayload.value ?? numberPayload.default,
        },
      };
  }

  throw new Error(`Invalid rule ${template.type}`);
};

// The convertRuleTemplateToPolicyRule will convert rule template to review policy rule for backend useage.
// Will throw exception if we don't implement the payload handler for specific type of rule.
export const convertRuleTemplateToPolicyRule = (
  rule: RuleTemplateV2
): SchemaPolicyRule => {
  const base: SchemaPolicyRule = {
    type: rule.type,
    level: rule.level,
    engine: rule.engine,
    comment: rule.comment ?? "",
  };
  if ((rule.componentList?.length ?? 0) === 0) {
    return base;
  }

  return mergeIndividualConfigAsRule(base, rule);
};

export const getRuleLocalizationKey = (type: string): string => {
  return type.split(".").join("-");
};

export const getRuleLocalization = (
  type: string
): { title: string; description: string } => {
  const key = getRuleLocalizationKey(type);

  const title = t(`sql-review.rule.${key}.title`);
  const description = t(`sql-review.rule.${key}.description`);

  return {
    title,
    description,
  };
};

export const ruleIsAvailableInSubscription = (
  ruleType: string,
  subscriptionPlan: PlanType
): boolean => {
  return true;
};

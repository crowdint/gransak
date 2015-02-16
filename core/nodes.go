package core

import (
	"strings"
)

type OperatorFunction func(re *GransakCore)

type Node struct {
	Name       string
	Nodes      []*Node
	Apply      OperatorFunction
	IsOperator bool
}

var Tree = &Node{
	Name: "Operators",
	Nodes: []*Node{
		&Node{
			Name: "or",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.template += "OR "
			},
		},
		&Node{
			Name: "and",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.template += "AND "
			},
		},
		&Node{
			Name: "eq",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("", "")
				re.replacePlaceholder("= " + re.valueholder)
			},
		},
		&Node{
			Name: "in",
			Apply: func(re *GransakCore) {
				re.appendField()
				items := []string{}

				numParams := re.param.parse("", "")

				for i := 1; i <= numParams; i++ {
					items = append(items, re.valueholder)
				}

				re.replacePlaceholder("IN (" + strings.Join(items, ",") + ")")
			},
		},
		&Node{
			Name: "matches",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("", "")
				re.replacePlaceholder("LIKE " + re.valueholder)
			},
		},
		&Node{
			Name: "cont",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("%", "%")
				re.replacePlaceholder("LIKE " + re.valueholder)
			},
			IsOperator: true,
			Nodes: []*Node{
				&Node{
					Name: "any",
					Apply: func(re *GransakCore) {
						field := re.getLastField()

						times := re.param.parse("%", "") - 1

						statement := field + " LIKE " + re.valueholder

						re.template += statement

						for i := 0; i < (times); i++ {
							re.template += " OR " + statement
						}
					},
				},
			},
		},
		&Node{
			Name: "lt",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("", "")
				re.replacePlaceholder("< " + re.valueholder)
			},
		},
		&Node{
			Name: "lteq",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("", "")
				re.replacePlaceholder("<= " + re.valueholder)
			},
		},
		&Node{
			Name: "gt",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("", "")
				re.replacePlaceholder("> " + re.valueholder)
			},
		},
		&Node{
			Name: "gteq",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("", "")
				re.replacePlaceholder(">= " + re.valueholder)
			},
		},
		&Node{
			Name: "start",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("", "%")
				re.replacePlaceholder("LIKE " + re.valueholder)
			},
		},
		&Node{
			Name: "end",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.param.parse("%", "")
				re.replacePlaceholder("LIKE " + re.valueholder)
			},
		},
		&Node{
			Name: "true",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.replacePlaceholder("= 't'")
			},
		},
		&Node{
			Name: "false",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.replacePlaceholder("= 'f'")
			},
		},
		&Node{
			Name: "present",
			Apply: func(re *GransakCore) {
				field := re.appendField()
				re.replacePlaceholder("IS NOT NULL AND " + field + " <> ''")
			},
		},
		&Node{
			Name: "blank",
			Apply: func(re *GransakCore) {
				field := re.appendField()
				re.replacePlaceholder("IS NULL OR " + field + " = ''")
			},
		},
		&Node{
			Name: "null",
			Apply: func(re *GransakCore) {
				re.appendField()
				re.replacePlaceholder("IS NULL")
			},
		},
		&Node{
			Name: "not",
			Nodes: []*Node{
				&Node{
					Name: "eq",
					Apply: func(re *GransakCore) {
						re.appendField()
						re.param.parse("", "")
						re.replacePlaceholder("<> " + re.valueholder)
					},
				},
				&Node{
					Name: "in",
					Apply: func(re *GransakCore) {
						re.appendField()
						items := []string{}

						numParams := re.param.parse("", "")

						for i := 1; i <= numParams; i++ {
							items = append(items, re.valueholder)
						}

						re.replacePlaceholder("NOT IN (" + strings.Join(items, ",") + ")")
					},
				},
				&Node{
					Name: "cont",
					Apply: func(re *GransakCore) {
						re.appendField()
						re.param.parse("%", "%")
						re.replacePlaceholder("NOT LIKE " + re.valueholder)
					},
					IsOperator: true,
					Nodes: []*Node{
						&Node{
							Name: "any",
							Apply: func(re *GransakCore) {
								field := re.getLastField()

								times := re.param.parse("%", "%") - 1

								statement := field + " NOT LIKE " + re.valueholder

								re.template += statement

								for i := 0; i < (times); i++ {
									re.template += " AND " + statement
								}
							},
						},
					},
				},
				&Node{
					Name: "start",
					Apply: func(re *GransakCore) {
						re.appendField()
						re.param.parse("", "%")
						re.replacePlaceholder("NOT LIKE " + re.valueholder)
					},
				},
				&Node{
					Name: "end",
					Apply: func(re *GransakCore) {
						re.appendField()
						re.param.parse("%", "")
						re.replacePlaceholder("NOT LIKE " + re.valueholder)
					},
				},
				&Node{
					Name: "true",
					Apply: func(re *GransakCore) {
						re.appendField()
						re.replacePlaceholder("<> 't'")
					},
				},
				&Node{
					Name: "false",
					Apply: func(re *GransakCore) {
						re.appendField()
						re.replacePlaceholder("<> 'f'")
					},
				},
				&Node{
					Name: "null",
					Apply: func(re *GransakCore) {
						re.appendField()
						re.replacePlaceholder("IS NOT NULL")
					},
				},
			},
		},
		&Node{
			Name: "does",
			Nodes: []*Node{
				&Node{
					Name: "not",
					Nodes: []*Node{
						&Node{
							Name: "match",
							Apply: func(re *GransakCore) {
								re.appendField()
								re.param.parse("", "")
								re.replacePlaceholder("NOT LIKE " + re.valueholder)
							},
						},
					},
				},
			},
		},
	},
}

package filter

type OperatorFunction func(re *GransakFilter)

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
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.template += "OR "
			},
		},
		&Node{
			Name: "and",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.template += "AND "
			},
		},
		&Node{
			Name: "eq",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("= " + re.getCorrectSqlFormat(re.valueholder))
			},
		},
		&Node{
			Name: "in",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("IN (" + re.valueholder + ")")
			},
		},
		&Node{
			Name: "matches",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("LIKE '" + re.valueholder + "'")
			},
		},
		&Node{
			Name: "cont",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("LIKE '%" + re.valueholder + "%'")
			},
			IsOperator: true,
			Nodes: []*Node{
				&Node{
					Name: "any",
					Apply: func(re *GransakFilter) {
						field := re.getLastField()

						values := re.param.parts
						times := len(values) - 1

						statement := field + " LIKE '%" + re.valueholder + "%'"

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
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("< " + re.getCorrectSqlFormat(re.valueholder))
			},
		},
		&Node{
			Name: "lteq",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("<= " + re.getCorrectSqlFormat(re.valueholder))
			},
		},
		&Node{
			Name: "gt",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("> " + re.getCorrectSqlFormat(re.valueholder))
			},
		},
		&Node{
			Name: "gteq",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder(">= " + re.getCorrectSqlFormat(re.valueholder))
			},
		},
		&Node{
			Name: "start",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("LIKE '" + re.valueholder + "%'")
			},
		},
		&Node{
			Name: "end",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("LIKE '%" + re.valueholder + "'")
			},
		},
		&Node{
			Name: "true",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("= 't'")
			},
		},
		&Node{
			Name: "false",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("= 'f'")
			},
		},
		&Node{
			Name: "present",
			Apply: func(re *GransakFilter) {
				field := re.appendField()
				re.replacePlaceholder("IS NOT NULL AND " + field + " <> ''")
			},
		},
		&Node{
			Name: "blank",
			Apply: func(re *GransakFilter) {
				field := re.appendField()
				re.replacePlaceholder("IS NULL OR " + re.placeholder)
				re.replacePlaceholder(field + " = ''")
			},
		},
		&Node{
			Name: "null",
			Apply: func(re *GransakFilter) {
				re.appendField()
				re.replacePlaceholder("IS NULL")
			},
		},
		&Node{
			Name: "not",
			Nodes: []*Node{
				&Node{
					Name: "eq",
					Apply: func(re *GransakFilter) {
						re.appendField()
						re.replacePlaceholder("<> " + re.getCorrectSqlFormat(re.valueholder))
					},
				},
				&Node{
					Name: "in",
					Apply: func(re *GransakFilter) {
						re.appendField()
						re.replacePlaceholder("NOT IN (" + re.valueholder + ")")
					},
				},
				&Node{
					Name: "cont",
					Apply: func(re *GransakFilter) {
						re.appendField()
						re.replacePlaceholder("NOT LIKE '%" + re.valueholder + "%'")
					},
					IsOperator: true,
					Nodes: []*Node{
						&Node{
							Name: "any",
							Apply: func(re *GransakFilter) {
								field := re.getLastField()

								values := re.param.parts
								times := len(values) - 1

								statement := field + " NOT LIKE '%" + re.valueholder + "%'"

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
					Apply: func(re *GransakFilter) {
						re.appendField()
						re.replacePlaceholder("NOT LIKE '" + re.valueholder + "%'")
					},
				},
				&Node{
					Name: "end",
					Apply: func(re *GransakFilter) {
						re.appendField()
						re.replacePlaceholder("NOT LIKE '%" + re.valueholder + "'")
					},
				},
				&Node{
					Name: "true",
					Apply: func(re *GransakFilter) {
						re.appendField()
						re.replacePlaceholder("<> 't'")
					},
				},
				&Node{
					Name: "false",
					Apply: func(re *GransakFilter) {
						re.appendField()
						re.replacePlaceholder("<> 'f'")
					},
				},
				&Node{
					Name: "null",
					Apply: func(re *GransakFilter) {
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
							Apply: func(re *GransakFilter) {
								re.appendField()
								re.replacePlaceholder("NOT LIKE '" + re.valueholder + "'")
							},
						},
					},
				},
			},
		},
	},
}

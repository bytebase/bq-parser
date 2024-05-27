// Code generated from bigquery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // bigquery
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by bigqueryParser.
type bigqueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by bigqueryParser#query_statement.
	VisitQuery_statement(ctx *Query_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#query_expr.
	VisitQuery_expr(ctx *Query_exprContext) interface{}

	// Visit a parse tree produced by bigqueryParser#select_statement.
	VisitSelect_statement(ctx *Select_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#from_statement.
	VisitFrom_statement(ctx *From_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#from_item.
	VisitFrom_item(ctx *From_itemContext) interface{}

	// Visit a parse tree produced by bigqueryParser#where_statement.
	VisitWhere_statement(ctx *Where_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#group_statement.
	VisitGroup_statement(ctx *Group_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#having_statement.
	VisitHaving_statement(ctx *Having_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#window_statement.
	VisitWindow_statement(ctx *Window_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#order_clause.
	VisitOrder_clause(ctx *Order_clauseContext) interface{}

	// Visit a parse tree produced by bigqueryParser#limit_clause.
	VisitLimit_clause(ctx *Limit_clauseContext) interface{}

	// Visit a parse tree produced by bigqueryParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}

	// Visit a parse tree produced by bigqueryParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by bigqueryParser#cast_expr.
	VisitCast_expr(ctx *Cast_exprContext) interface{}

	// Visit a parse tree produced by bigqueryParser#column_expr.
	VisitColumn_expr(ctx *Column_exprContext) interface{}

	// Visit a parse tree produced by bigqueryParser#except_statement.
	VisitExcept_statement(ctx *Except_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#replace_statement.
	VisitReplace_statement(ctx *Replace_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#join_type.
	VisitJoin_type(ctx *Join_typeContext) interface{}

	// Visit a parse tree produced by bigqueryParser#on_clause.
	VisitOn_clause(ctx *On_clauseContext) interface{}

	// Visit a parse tree produced by bigqueryParser#set_op.
	VisitSet_op(ctx *Set_opContext) interface{}

	// Visit a parse tree produced by bigqueryParser#using_clause.
	VisitUsing_clause(ctx *Using_clauseContext) interface{}

	// Visit a parse tree produced by bigqueryParser#field_path.
	VisitField_path(ctx *Field_pathContext) interface{}

	// Visit a parse tree produced by bigqueryParser#sstruct.
	VisitSstruct(ctx *SstructContext) interface{}

	// Visit a parse tree produced by bigqueryParser#array_expr.
	VisitArray_expr(ctx *Array_exprContext) interface{}

	// Visit a parse tree produced by bigqueryParser#array_path.
	VisitArray_path(ctx *Array_pathContext) interface{}

	// Visit a parse tree produced by bigqueryParser#bool_expression.
	VisitBool_expression(ctx *Bool_expressionContext) interface{}

	// Visit a parse tree produced by bigqueryParser#window_name.
	VisitWindow_name(ctx *Window_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#window_definition.
	VisitWindow_definition(ctx *Window_definitionContext) interface{}

	// Visit a parse tree produced by bigqueryParser#count.
	VisitCount(ctx *CountContext) interface{}

	// Visit a parse tree produced by bigqueryParser#skip_rows.
	VisitSkip_rows(ctx *Skip_rowsContext) interface{}

	// Visit a parse tree produced by bigqueryParser#with_statement.
	VisitWith_statement(ctx *With_statementContext) interface{}

	// Visit a parse tree produced by bigqueryParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#alias_name.
	VisitAlias_name(ctx *Alias_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#array_name.
	VisitArray_name(ctx *Array_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#cte_name.
	VisitCte_name(ctx *Cte_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#dataset_name.
	VisitDataset_name(ctx *Dataset_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#datatype_name.
	VisitDatatype_name(ctx *Datatype_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#function_name.
	VisitFunction_name(ctx *Function_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#join_name.
	VisitJoin_name(ctx *Join_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#member_name.
	VisitMember_name(ctx *Member_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#project_name.
	VisitProject_name(ctx *Project_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#struct_name.
	VisitStruct_name(ctx *Struct_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by bigqueryParser#table_expr.
	VisitTable_expr(ctx *Table_exprContext) interface{}

	// Visit a parse tree produced by bigqueryParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by bigqueryParser#integer_type.
	VisitInteger_type(ctx *Integer_typeContext) interface{}

	// Visit a parse tree produced by bigqueryParser#float_type.
	VisitFloat_type(ctx *Float_typeContext) interface{}

	// Visit a parse tree produced by bigqueryParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by bigqueryParser#quoted_string.
	VisitQuoted_string(ctx *Quoted_stringContext) interface{}

	// Visit a parse tree produced by bigqueryParser#triple_quoted_string.
	VisitTriple_quoted_string(ctx *Triple_quoted_stringContext) interface{}

	// Visit a parse tree produced by bigqueryParser#raw_string.
	VisitRaw_string(ctx *Raw_stringContext) interface{}

	// Visit a parse tree produced by bigqueryParser#byte_string.
	VisitByte_string(ctx *Byte_stringContext) interface{}

	// Visit a parse tree produced by bigqueryParser#raw_byte_string.
	VisitRaw_byte_string(ctx *Raw_byte_stringContext) interface{}

	// Visit a parse tree produced by bigqueryParser#special_string.
	VisitSpecial_string(ctx *Special_stringContext) interface{}

	// Visit a parse tree produced by bigqueryParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}
}

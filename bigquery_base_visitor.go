// Code generated from bigquery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // bigquery
import "github.com/antlr4-go/antlr/v4"

type BasebigqueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasebigqueryVisitor) VisitQuery_statement(ctx *Query_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitQuery_expr(ctx *Query_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitSelect_statement(ctx *Select_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitFrom_statement(ctx *From_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitFrom_item(ctx *From_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitWhere_statement(ctx *Where_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitGroup_statement(ctx *Group_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitHaving_statement(ctx *Having_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitWindow_statement(ctx *Window_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitOrder_clause(ctx *Order_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitLimit_clause(ctx *Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitCast_expr(ctx *Cast_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitColumn_expr(ctx *Column_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitExcept_statement(ctx *Except_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitReplace_statement(ctx *Replace_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitJoin_type(ctx *Join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitOn_clause(ctx *On_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitSet_op(ctx *Set_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitField_path(ctx *Field_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitSstruct(ctx *SstructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitArray_expr(ctx *Array_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitArray_path(ctx *Array_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitBool_expression(ctx *Bool_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitWindow_name(ctx *Window_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitWindow_definition(ctx *Window_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitCount(ctx *CountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitSkip_rows(ctx *Skip_rowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitWith_statement(ctx *With_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitAlias_name(ctx *Alias_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitArray_name(ctx *Array_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitCte_name(ctx *Cte_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitDataset_name(ctx *Dataset_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitDatatype_name(ctx *Datatype_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitJoin_name(ctx *Join_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitMember_name(ctx *Member_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitProject_name(ctx *Project_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitStruct_name(ctx *Struct_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitTable_expr(ctx *Table_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitInteger_type(ctx *Integer_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitFloat_type(ctx *Float_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitQuoted_string(ctx *Quoted_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitTriple_quoted_string(ctx *Triple_quoted_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitRaw_string(ctx *Raw_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitByte_string(ctx *Byte_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitRaw_byte_string(ctx *Raw_byte_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitSpecial_string(ctx *Special_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebigqueryVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

// Generated from Protobuf3.g4 by ANTLR 4.7.

package pb // Protobuf3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProtobuf3Listener is a complete listener for a parse tree produced by Protobuf3Parser.
type BaseProtobuf3Listener struct{}

var _ Protobuf3Listener = &BaseProtobuf3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProtobuf3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProtobuf3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProtobuf3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProtobuf3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *BaseProtobuf3Listener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *BaseProtobuf3Listener) ExitProto(ctx *ProtoContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *BaseProtobuf3Listener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BaseProtobuf3Listener) ExitSyntax(ctx *SyntaxContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseProtobuf3Listener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseProtobuf3Listener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *BaseProtobuf3Listener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *BaseProtobuf3Listener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterOption is called when production option is entered.
func (s *BaseProtobuf3Listener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseProtobuf3Listener) ExitOption(ctx *OptionContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *BaseProtobuf3Listener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *BaseProtobuf3Listener) ExitOptionName(ctx *OptionNameContext) {}

// EnterTopLevelDef is called when production topLevelDef is entered.
func (s *BaseProtobuf3Listener) EnterTopLevelDef(ctx *TopLevelDefContext) {}

// ExitTopLevelDef is called when production topLevelDef is exited.
func (s *BaseProtobuf3Listener) ExitTopLevelDef(ctx *TopLevelDefContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseProtobuf3Listener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseProtobuf3Listener) ExitMessage(ctx *MessageContext) {}

// EnterMessageBody is called when production messageBody is entered.
func (s *BaseProtobuf3Listener) EnterMessageBody(ctx *MessageBodyContext) {}

// ExitMessageBody is called when production messageBody is exited.
func (s *BaseProtobuf3Listener) ExitMessageBody(ctx *MessageBodyContext) {}

// EnterEnumDefinition is called when production enumDefinition is entered.
func (s *BaseProtobuf3Listener) EnterEnumDefinition(ctx *EnumDefinitionContext) {}

// ExitEnumDefinition is called when production enumDefinition is exited.
func (s *BaseProtobuf3Listener) ExitEnumDefinition(ctx *EnumDefinitionContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseProtobuf3Listener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseProtobuf3Listener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BaseProtobuf3Listener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BaseProtobuf3Listener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterEnumValueOption is called when production enumValueOption is entered.
func (s *BaseProtobuf3Listener) EnterEnumValueOption(ctx *EnumValueOptionContext) {}

// ExitEnumValueOption is called when production enumValueOption is exited.
func (s *BaseProtobuf3Listener) ExitEnumValueOption(ctx *EnumValueOptionContext) {}

// EnterService is called when production service is entered.
func (s *BaseProtobuf3Listener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseProtobuf3Listener) ExitService(ctx *ServiceContext) {}

// EnterRpc is called when production rpc is entered.
func (s *BaseProtobuf3Listener) EnterRpc(ctx *RpcContext) {}

// ExitRpc is called when production rpc is exited.
func (s *BaseProtobuf3Listener) ExitRpc(ctx *RpcContext) {}

// EnterReserved is called when production reserved is entered.
func (s *BaseProtobuf3Listener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *BaseProtobuf3Listener) ExitReserved(ctx *ReservedContext) {}

// EnterPbranges is called when production pbranges is entered.
func (s *BaseProtobuf3Listener) EnterPbranges(ctx *PbrangesContext) {}

// ExitPbranges is called when production pbranges is exited.
func (s *BaseProtobuf3Listener) ExitPbranges(ctx *PbrangesContext) {}

// EnterPbrange is called when production pbrange is entered.
func (s *BaseProtobuf3Listener) EnterPbrange(ctx *PbrangeContext) {}

// ExitPbrange is called when production pbrange is exited.
func (s *BaseProtobuf3Listener) ExitPbrange(ctx *PbrangeContext) {}

// EnterFieldNames is called when production fieldNames is entered.
func (s *BaseProtobuf3Listener) EnterFieldNames(ctx *FieldNamesContext) {}

// ExitFieldNames is called when production fieldNames is exited.
func (s *BaseProtobuf3Listener) ExitFieldNames(ctx *FieldNamesContext) {}

// EnterPbtype is called when production pbtype is entered.
func (s *BaseProtobuf3Listener) EnterPbtype(ctx *PbtypeContext) {}

// ExitPbtype is called when production pbtype is exited.
func (s *BaseProtobuf3Listener) ExitPbtype(ctx *PbtypeContext) {}

// EnterFieldNumber is called when production fieldNumber is entered.
func (s *BaseProtobuf3Listener) EnterFieldNumber(ctx *FieldNumberContext) {}

// ExitFieldNumber is called when production fieldNumber is exited.
func (s *BaseProtobuf3Listener) ExitFieldNumber(ctx *FieldNumberContext) {}

// EnterField is called when production field is entered.
func (s *BaseProtobuf3Listener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseProtobuf3Listener) ExitField(ctx *FieldContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BaseProtobuf3Listener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BaseProtobuf3Listener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterFieldOption is called when production fieldOption is entered.
func (s *BaseProtobuf3Listener) EnterFieldOption(ctx *FieldOptionContext) {}

// ExitFieldOption is called when production fieldOption is exited.
func (s *BaseProtobuf3Listener) ExitFieldOption(ctx *FieldOptionContext) {}

// EnterOneof is called when production oneof is entered.
func (s *BaseProtobuf3Listener) EnterOneof(ctx *OneofContext) {}

// ExitOneof is called when production oneof is exited.
func (s *BaseProtobuf3Listener) ExitOneof(ctx *OneofContext) {}

// EnterOneofField is called when production oneofField is entered.
func (s *BaseProtobuf3Listener) EnterOneofField(ctx *OneofFieldContext) {}

// ExitOneofField is called when production oneofField is exited.
func (s *BaseProtobuf3Listener) ExitOneofField(ctx *OneofFieldContext) {}

// EnterMapField is called when production mapField is entered.
func (s *BaseProtobuf3Listener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *BaseProtobuf3Listener) ExitMapField(ctx *MapFieldContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BaseProtobuf3Listener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BaseProtobuf3Listener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *BaseProtobuf3Listener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *BaseProtobuf3Listener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterMessageName is called when production messageName is entered.
func (s *BaseProtobuf3Listener) EnterMessageName(ctx *MessageNameContext) {}

// ExitMessageName is called when production messageName is exited.
func (s *BaseProtobuf3Listener) ExitMessageName(ctx *MessageNameContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *BaseProtobuf3Listener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *BaseProtobuf3Listener) ExitEnumName(ctx *EnumNameContext) {}

// EnterMessageOrEnumName is called when production messageOrEnumName is entered.
func (s *BaseProtobuf3Listener) EnterMessageOrEnumName(ctx *MessageOrEnumNameContext) {}

// ExitMessageOrEnumName is called when production messageOrEnumName is exited.
func (s *BaseProtobuf3Listener) ExitMessageOrEnumName(ctx *MessageOrEnumNameContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseProtobuf3Listener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseProtobuf3Listener) ExitFieldName(ctx *FieldNameContext) {}

// EnterOneofName is called when production oneofName is entered.
func (s *BaseProtobuf3Listener) EnterOneofName(ctx *OneofNameContext) {}

// ExitOneofName is called when production oneofName is exited.
func (s *BaseProtobuf3Listener) ExitOneofName(ctx *OneofNameContext) {}

// EnterMapName is called when production mapName is entered.
func (s *BaseProtobuf3Listener) EnterMapName(ctx *MapNameContext) {}

// ExitMapName is called when production mapName is exited.
func (s *BaseProtobuf3Listener) ExitMapName(ctx *MapNameContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *BaseProtobuf3Listener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *BaseProtobuf3Listener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterRpcName is called when production rpcName is entered.
func (s *BaseProtobuf3Listener) EnterRpcName(ctx *RpcNameContext) {}

// ExitRpcName is called when production rpcName is exited.
func (s *BaseProtobuf3Listener) ExitRpcName(ctx *RpcNameContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *BaseProtobuf3Listener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *BaseProtobuf3Listener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterMessageOrEnumType is called when production messageOrEnumType is entered.
func (s *BaseProtobuf3Listener) EnterMessageOrEnumType(ctx *MessageOrEnumTypeContext) {}

// ExitMessageOrEnumType is called when production messageOrEnumType is exited.
func (s *BaseProtobuf3Listener) ExitMessageOrEnumType(ctx *MessageOrEnumTypeContext) {}

// EnterEmptyStatement is called when production emptyStatement is entered.
func (s *BaseProtobuf3Listener) EnterEmptyStatement(ctx *EmptyStatementContext) {}

// ExitEmptyStatement is called when production emptyStatement is exited.
func (s *BaseProtobuf3Listener) ExitEmptyStatement(ctx *EmptyStatementContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseProtobuf3Listener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseProtobuf3Listener) ExitConstant(ctx *ConstantContext) {}

// Generated from Protobuf3.g4 by ANTLR 4.7.

package pb // Protobuf3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Protobuf3Listener is a complete listener for a parse tree produced by Protobuf3Parser.
type Protobuf3Listener interface {
	antlr.ParseTreeListener

	// EnterProto is called when entering the proto production.
	EnterProto(c *ProtoContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterPackageStatement is called when entering the packageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterOptionName is called when entering the optionName production.
	EnterOptionName(c *OptionNameContext)

	// EnterTopLevelDef is called when entering the topLevelDef production.
	EnterTopLevelDef(c *TopLevelDefContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterMessageBody is called when entering the messageBody production.
	EnterMessageBody(c *MessageBodyContext)

	// EnterEnumDefinition is called when entering the enumDefinition production.
	EnterEnumDefinition(c *EnumDefinitionContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterEnumValueOption is called when entering the enumValueOption production.
	EnterEnumValueOption(c *EnumValueOptionContext)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterRpc is called when entering the rpc production.
	EnterRpc(c *RpcContext)

	// EnterReserved is called when entering the reserved production.
	EnterReserved(c *ReservedContext)

	// EnterPbranges is called when entering the pbranges production.
	EnterPbranges(c *PbrangesContext)

	// EnterPbrange is called when entering the pbrange production.
	EnterPbrange(c *PbrangeContext)

	// EnterFieldNames is called when entering the fieldNames production.
	EnterFieldNames(c *FieldNamesContext)

	// EnterPbtype is called when entering the pbtype production.
	EnterPbtype(c *PbtypeContext)

	// EnterFieldNumber is called when entering the fieldNumber production.
	EnterFieldNumber(c *FieldNumberContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldOptions is called when entering the fieldOptions production.
	EnterFieldOptions(c *FieldOptionsContext)

	// EnterFieldOption is called when entering the fieldOption production.
	EnterFieldOption(c *FieldOptionContext)

	// EnterOneof is called when entering the oneof production.
	EnterOneof(c *OneofContext)

	// EnterOneofField is called when entering the oneofField production.
	EnterOneofField(c *OneofFieldContext)

	// EnterMapField is called when entering the mapField production.
	EnterMapField(c *MapFieldContext)

	// EnterKeyType is called when entering the keyType production.
	EnterKeyType(c *KeyTypeContext)

	// EnterFullIdent is called when entering the fullIdent production.
	EnterFullIdent(c *FullIdentContext)

	// EnterMessageName is called when entering the messageName production.
	EnterMessageName(c *MessageNameContext)

	// EnterEnumName is called when entering the enumName production.
	EnterEnumName(c *EnumNameContext)

	// EnterMessageOrEnumName is called when entering the messageOrEnumName production.
	EnterMessageOrEnumName(c *MessageOrEnumNameContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterOneofName is called when entering the oneofName production.
	EnterOneofName(c *OneofNameContext)

	// EnterMapName is called when entering the mapName production.
	EnterMapName(c *MapNameContext)

	// EnterServiceName is called when entering the serviceName production.
	EnterServiceName(c *ServiceNameContext)

	// EnterRpcName is called when entering the rpcName production.
	EnterRpcName(c *RpcNameContext)

	// EnterMessageType is called when entering the messageType production.
	EnterMessageType(c *MessageTypeContext)

	// EnterMessageOrEnumType is called when entering the messageOrEnumType production.
	EnterMessageOrEnumType(c *MessageOrEnumTypeContext)

	// EnterEmptyStatement is called when entering the emptyStatement production.
	EnterEmptyStatement(c *EmptyStatementContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// ExitProto is called when exiting the proto production.
	ExitProto(c *ProtoContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitPackageStatement is called when exiting the packageStatement production.
	ExitPackageStatement(c *PackageStatementContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitOptionName is called when exiting the optionName production.
	ExitOptionName(c *OptionNameContext)

	// ExitTopLevelDef is called when exiting the topLevelDef production.
	ExitTopLevelDef(c *TopLevelDefContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitMessageBody is called when exiting the messageBody production.
	ExitMessageBody(c *MessageBodyContext)

	// ExitEnumDefinition is called when exiting the enumDefinition production.
	ExitEnumDefinition(c *EnumDefinitionContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitEnumValueOption is called when exiting the enumValueOption production.
	ExitEnumValueOption(c *EnumValueOptionContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitRpc is called when exiting the rpc production.
	ExitRpc(c *RpcContext)

	// ExitReserved is called when exiting the reserved production.
	ExitReserved(c *ReservedContext)

	// ExitPbranges is called when exiting the pbranges production.
	ExitPbranges(c *PbrangesContext)

	// ExitPbrange is called when exiting the pbrange production.
	ExitPbrange(c *PbrangeContext)

	// ExitFieldNames is called when exiting the fieldNames production.
	ExitFieldNames(c *FieldNamesContext)

	// ExitPbtype is called when exiting the pbtype production.
	ExitPbtype(c *PbtypeContext)

	// ExitFieldNumber is called when exiting the fieldNumber production.
	ExitFieldNumber(c *FieldNumberContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldOptions is called when exiting the fieldOptions production.
	ExitFieldOptions(c *FieldOptionsContext)

	// ExitFieldOption is called when exiting the fieldOption production.
	ExitFieldOption(c *FieldOptionContext)

	// ExitOneof is called when exiting the oneof production.
	ExitOneof(c *OneofContext)

	// ExitOneofField is called when exiting the oneofField production.
	ExitOneofField(c *OneofFieldContext)

	// ExitMapField is called when exiting the mapField production.
	ExitMapField(c *MapFieldContext)

	// ExitKeyType is called when exiting the keyType production.
	ExitKeyType(c *KeyTypeContext)

	// ExitFullIdent is called when exiting the fullIdent production.
	ExitFullIdent(c *FullIdentContext)

	// ExitMessageName is called when exiting the messageName production.
	ExitMessageName(c *MessageNameContext)

	// ExitEnumName is called when exiting the enumName production.
	ExitEnumName(c *EnumNameContext)

	// ExitMessageOrEnumName is called when exiting the messageOrEnumName production.
	ExitMessageOrEnumName(c *MessageOrEnumNameContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitOneofName is called when exiting the oneofName production.
	ExitOneofName(c *OneofNameContext)

	// ExitMapName is called when exiting the mapName production.
	ExitMapName(c *MapNameContext)

	// ExitServiceName is called when exiting the serviceName production.
	ExitServiceName(c *ServiceNameContext)

	// ExitRpcName is called when exiting the rpcName production.
	ExitRpcName(c *RpcNameContext)

	// ExitMessageType is called when exiting the messageType production.
	ExitMessageType(c *MessageTypeContext)

	// ExitMessageOrEnumType is called when exiting the messageOrEnumType production.
	ExitMessageOrEnumType(c *MessageOrEnumTypeContext)

	// ExitEmptyStatement is called when exiting the emptyStatement production.
	ExitEmptyStatement(c *EmptyStatementContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)
}

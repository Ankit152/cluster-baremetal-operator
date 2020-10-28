// Code generated by "stringer -type=BuiltinPluginType"; DO NOT EDIT.

package builtinhelpers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[AnnotationsTransformer-1]
	_ = x[ConfigMapGenerator-2]
	_ = x[HashTransformer-3]
	_ = x[ImageTagTransformer-4]
	_ = x[LabelTransformer-5]
	_ = x[LegacyOrderTransformer-6]
	_ = x[NamespaceTransformer-7]
	_ = x[PatchJson6902Transformer-8]
	_ = x[PatchStrategicMergeTransformer-9]
	_ = x[PatchTransformer-10]
	_ = x[PrefixSuffixTransformer-11]
	_ = x[ReplicaCountTransformer-12]
	_ = x[SecretGenerator-13]
	_ = x[ValueAddTransformer-14]
}

const _BuiltinPluginType_name = "UnknownAnnotationsTransformerConfigMapGeneratorHashTransformerImageTagTransformerLabelTransformerLegacyOrderTransformerNamespaceTransformerPatchJson6902TransformerPatchStrategicMergeTransformerPatchTransformerPrefixSuffixTransformerReplicaCountTransformerSecretGeneratorValueAddTransformer"

var _BuiltinPluginType_index = [...]uint16{0, 7, 29, 47, 62, 81, 97, 119, 139, 163, 193, 209, 232, 255, 270, 289}

func (i BuiltinPluginType) String() string {
	if i < 0 || i >= BuiltinPluginType(len(_BuiltinPluginType_index)-1) {
		return "BuiltinPluginType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BuiltinPluginType_name[_BuiltinPluginType_index[i]:_BuiltinPluginType_index[i+1]]
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AttackPattern holds the schema definition for the AttackPattern entity.
type AttackPattern struct {
	ent.Schema
}

// Fields of the AttackPattern.
func (AttackPattern) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("type"),
		field.String("name"),
		field.String("aliases").Nillable().Optional(),
		field.String("description").Nillable().Optional(),
		field.String("kill_chain_phases").Nillable().Optional(),
		field.String("spec_version").Nillable().Optional(),
		field.String("created_by_ref").Nillable().Optional(),
		field.String("labels").Nillable().Optional(),
		field.Time("created").Nillable().Optional(),
		field.Time("modified").Nillable().Optional(),
		field.Bool("revoked").Nillable().Optional(),
		field.String("confidence").Nillable().Optional(),
		field.String("lang").Nillable().Optional(),
		field.String("external_references").Nillable().Optional(),
		field.String("object_marking_refs").Nillable().Optional(),
		field.String("granular_markings").Nillable().Optional(),
		field.String("extensions").Nillable().Optional(),
		field.String("x_mitre_platforms").Nillable().Optional(),
		field.String("x_mitre_domains").Nillable().Optional(),
		field.String("x_mitre_detection").Nillable().Optional(),
		field.String("x_mitre_version").Nillable().Optional(),
		field.String("x_mitre_modified_by_ref").Nillable().Optional(),
		field.String("x_mitre_defense_bypassed").Nillable().Optional(),
		field.Bool("x_mitre_is_subtechnique").Nillable().Optional(),
		field.Bool("x_mitre_deprecated").Nillable().Optional(),
		field.Bool("x_mitre_remote_support").Nillable().Optional(),
		field.String("x_mitre_data_sources").Nillable().Optional(),
		field.String("x_mitre_attack_spec_version").Nillable().Optional(),
		field.String("x_mitre_system_requirements").Nillable().Optional(),
		field.String("x_mitre_contributors").Nillable().Optional(),
		field.String("x_mitre_permissions_required").Nillable().Optional(),
	}
}

// Edges of the AttackPattern.
func (AttackPattern) Edges() []ent.Edge {
	return nil
}

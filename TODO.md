## To Do
### ``assignment.go``
(line 225) (kenjones): Add support for the evaluation of ConstraintClause


### ``flattener.go``
(line 14) (kenjones): Switch to reflect if possible in the future (if simple)


### ``parser.go``
(line 84) (kenjones): Add hooks as method parameter

(line 139) (kenjones): Does dropping the Imports list really have any impact?


### ``service_template.go``
(line 21) (kenjones): Implement ImportDefinition as it is not always going to be a simple

(line 187) (kenjones): assume the requirement has a node specified, otherwise need to use the


### ``topology.go``
(line 62) (kenjones): Add support for Groups


### ``tosca_namespace_alias.go``
(line 43) (kenjones): Leverage https://github.com/blang/semver to provide Version implementation details.

(line 45) Version.GetMajor

(line 52) Version.GetMinor

(line 59) Version.GetFixVersion

(line 66) Version.GetQualifier

(line 73) Version.GetBuildVersion


### ``tosca_reusable_modeling_definitions.go``
(line 19) : Implement ArtifactDefinition struct

(line 22) : Implement NodeFilter struct

(line 70) : Implement ArtifactType struct


package model

import (
	"fmt"
	"github.com/hellgate75/k8s-deploy/utils"
)

type Response struct {
	Status    int          `yaml:"status" json:"status" xml:"status"`
	Message   string       `yaml:"message" json:"message" xml:"message"`
	Reference ApiReference `yaml:"reference" json:"reference" xml:"reference"`
	Data      interface{}  `yaml:"data" json:"data" xml:"data"`
}

type ApiReferenceItem struct {
	Name string `yaml:"name" json:"name" xml:"name"`
	Url  string `yaml:"url" json:"url" xml:"url"`
}

type ApiReference struct {
	CurrentUrl    string             `yaml:"_self" json:"_self" xml:"self"`
	CurrentMethod string             `yaml:"_method" json:"_method" xml:"method"`
	Urls          []ApiReferenceItem `yaml:"urls" json:"urls" xml:"url"`
}

type ResourceType string
type Aggregator string
type Oper string
type State string

const (
	ResourceTypeRepositories ResourceType = "k8srepo"
	ResourceTypeCharts       ResourceType = "charts"
	ResourceTypeChart        ResourceType = "chart"
	ResourceTypeKubeFiles    ResourceType = "kubefiles"
	ResourceTypeKubeFile     ResourceType = "kubefile"
	ResourceTypeDeploys      ResourceType = "deploys"
	ResourceTypeDeploy       ResourceType = "deploy"
	ResourceTypeJobs         ResourceType = "jobs"
	ResourceTypeJob          ResourceType = "job"
	ResourceTypeProjects     ResourceType = "projects"
	ResourceTypeProject      ResourceType = "project"

	AggregatorEq      Aggregator = "eq"
	AggregatorIn      Aggregator = "in"
	AggregatorLike    Aggregator = "like"
	AggregatorNeq     Aggregator = "neq"
	AggregatorNotIn   Aggregator = "nin"
	AggregatorNotLike Aggregator = "nlike"
	AggregatorNot     Aggregator = "not" //for bool

	OperOr   Oper = "or"
	OperAnd  Oper = "and"
	OperNor  Oper = "nor"
	OperNAnd Oper = "nand"

	StateCreated  State = "created"
	StateError    State = "error"
	StateReady    State = "ready"
	StateRunning  State = "running"
	StateComplete State = "complete"
	StateFailed   State = "failed"
	StateRollback State = "rolled-back"
	StateDeleting State = "deleting"
	StateDeleted  State = "deleted"
	StatePurging  State = "purging"
	StatePutged   State = "purged"
)

type QueryItem struct {
	Key        string     `yaml:"key" json:"key" xml:"key"`
	Value      string     `yaml:"value" json:"value" xml:"value"`
	Aggregator Aggregator `yaml:"aggregator" json:"aggregator" xml:"aggregator"`
}

func (q *QueryItem) ToJson() (string, error) {
	d, err := utils.StructureToJson(q)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (q *QueryItem) String() string {
	s, err := q.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (q *QueryItem) FromJson(d string) error {
	return utils.JsonToStructure(d, q)
}

func (q *QueryItem) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, q)
}

type Query struct {
	Items []QueryItem `yaml:"items" json:"items" xml:"item"`
	Oper  Oper        `yaml:"oper" json:"oper" xml:"oper"`
}

func (q *Query) ToJson() (string, error) {
	d, err := utils.StructureToJson(q)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (req *Request) ToJson() (string, error) {
	d, err := utils.StructureToJson(req)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (req *Request) String() string {
	s, err := req.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (req *Request) FromJson(d string) error {
	return utils.JsonToStructure(d, req)
}

func (req *Request) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, req)
}

type Version struct {
	Id    string `yaml:"id" json:"id" xml:"id"`
	Name  string `yaml:"name" json:"name" xml:"name"`
	State State  `yaml:"state" json:"state" xml:"state"`
}

func (ver *Version) ToJson() (string, error) {
	d, err := utils.StructureToJson(ver)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (ver *Version) String() string {
	s, err := ver.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (ver *Version) FromJson(d string) error {
	return utils.JsonToStructure(d, ver)
}

func (ver *Version) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, ver)
}

type VersionDetails struct {
	VersionId   string            `yaml:"versionId" json:"versionId" xml:"version-id"`
	VersionName string            `yaml:"versionName" json:"versionName" xml:"version-name"`
	Keys        map[string]string `yaml:"keys" json:"keys" xml:"key-entry"`
}

func (ver *VersionDetails) ToJson() (string, error) {
	d, err := utils.StructureToJson(ver)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (ver *VersionDetails) String() string {
	s, err := ver.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (ver *VersionDetails) FromJson(d string) error {
	return utils.JsonToStructure(d, ver)
}

func (ver *VersionDetails) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, ver)
}

type Chart struct {
	Id       string    `yaml:"id" json:"id" xml:"id"`
	Name     string    `yaml:"name" json:"name" xml:"name"`
	Versions []Version `yaml:"versions" json:"versions" xml:"version"`
	State    State     `yaml:"state" json:"state" xml:"state"`
}

func (ch *Chart) ToJson() (string, error) {
	d, err := utils.StructureToJson(ch)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (ch *Chart) String() string {
	s, err := ch.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (ch *Chart) FromJson(d string) error {
	return utils.JsonToStructure(d, ch)
}

func (ch *Chart) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, ch)
}

type KubeFile struct {
	Id       string    `yaml:"id" json:"id" xml:"id"`
	Name     string    `yaml:"name" json:"name" xml:"name"`
	Versions []Version `yaml:"versions" json:"versions" xml:"version"`
	State    State     `yaml:"state" json:"state" xml:"state"`
}

func (kf *KubeFile) ToJson() (string, error) {
	d, err := utils.StructureToJson(kf)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (kf *KubeFile) String() string {
	s, err := kf.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (kf *KubeFile) FromJson(d string) error {
	return utils.JsonToStructure(d, kf)
}

func (kf *KubeFile) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, kf)
}

type Repository struct {
	Id        string     `yaml:"id" json:"id" xml:"id"`
	Name      string     `yaml:"name" json:"name" xml:"name"`
	Charts    []Chart    `yaml:"charts" json:"charts" xml:"chart"`
	KubeFiles []KubeFile `yaml:"kubefiles" json:"kubefiles" xml:"kubefile"`
	State     State      `yaml:"state" json:"state" xml:"state"`
}

func (r *Repository) ToJson() (string, error) {
	d, err := utils.StructureToJson(r)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (r *Repository) String() string {
	s, err := r.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (r *Repository) FromJson(d string) error {
	return utils.JsonToStructure(d, r)
}

func (r *Repository) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, r)
}

type ProjectChart struct {
	Id      string `yaml:"id" json:"id" xml:"id"`
	Name    string `yaml:"name" json:"name" xml:"name"`
	Version string `yaml:"version" json:"version" xml:"version"`
	State   State  `yaml:"state" json:"state" xml:"state"`
}

func (pch *ProjectChart) ToJson() (string, error) {
	d, err := utils.StructureToJson(pch)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (pch *ProjectChart) String() string {
	s, err := pch.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (pch *ProjectChart) FromJson(d string) error {
	return utils.JsonToStructure(d, pch)
}

func (pch *ProjectChart) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, pch)
}

type ProjectKubeFile struct {
	Id      string `yaml:"id" json:"id" xml:"id"`
	Name    string `yaml:"name" json:"name" xml:"name"`
	Version string `yaml:"version" json:"version" xml:"version"`
	State   State  `yaml:"state" json:"state" xml:"state"`
}

func (pkf *ProjectKubeFile) ToJson() (string, error) {
	d, err := utils.StructureToJson(pkf)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (pkf *ProjectKubeFile) String() string {
	s, err := pkf.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (pkf *ProjectKubeFile) FromJson(d string) error {
	return utils.JsonToStructure(d, pkf)
}

func (pkf *ProjectKubeFile) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, pkf)
}

type ProjectVersion struct {
	Version   string            `yaml:"version" json:"version" xml:"version"`
	Charts    []ProjectChart    `yaml:"charts" json:"charts" xml:"chart"`
	KubeFiles []ProjectKubeFile `yaml:"kubefiles" json:"kubefiles" xml:"kubefile"`
	Variables []Variable        `yaml:"variables" json:"variables" xml:"variables"`
	State     State             `yaml:"state" json:"state" xml:"state"`
}

func (pv *ProjectVersion) ToJson() (string, error) {
	d, err := utils.StructureToJson(pv)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (pv *ProjectVersion) String() string {
	s, err := pv.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (pv *ProjectVersion) FromJson(d string) error {
	return utils.JsonToStructure(d, pv)
}

func (pv *ProjectVersion) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, pv)
}

type Project struct {
	Id       string           `yaml:"id" json:"id" xml:"id"`
	Name     string           `yaml:"name" json:"name" xml:"name"`
	Version  string           `yaml:"version" json:"version" xml:"version"`
	Versions []ProjectVersion `yaml:"versions" json:"versions" xml:"version"`
	State    State            `yaml:"state" json:"state" xml:"state"`
	ReadOnly bool             `yaml:"readOnly" json:"readOnly" xml:"read-only"`
}

func (p *Project) ToJson() (string, error) {
	d, err := utils.StructureToJson(p)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (p *Project) String() string {
	s, err := p.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (p *Project) FromJson(d string) error {
	return utils.JsonToStructure(d, p)
}

func (p *Project) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, p)
}

type VariableRule struct {
	Id        string `yaml:"id" json:"id" xml:"id"`
	Name      string `yaml:"name" json:"name" xml:"name"`
	ValidIf   string `yaml:"validif" json:"validif" xml:"valid-if"`
	InvalidIf string `yaml:"invalidif" json:"invalidif" xml:"invalid-if"`
}

func (vr *VariableRule) ToJson() (string, error) {
	d, err := utils.StructureToJson(vr)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (vr *VariableRule) String() string {
	s, err := vr.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (vr *VariableRule) FromJson(d string) error {
	return utils.JsonToStructure(d, vr)
}

func (vr *VariableRule) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, vr)
}

type Variable struct {
	Id      string         `yaml:"id" json:"id" xml:"id"`
	Name    string         `yaml:"name" json:"name" xml:"name"`
	Default interface{}    `yaml:"default" json:"default" xml:"default"`
	Rules   []VariableRule `yaml:"rules" json:"rules" xml:"rules"`
}

func (v *Variable) ToJson() (string, error) {
	d, err := utils.StructureToJson(v)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (v *Variable) String() string {
	s, err := v.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (v *Variable) FromJson(d string) error {
	return utils.JsonToStructure(d, v)
}

func (v *Variable) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, v)
}

type Param struct {
	Name  string      `yaml:"name" json:"name" xml:"name"`
	Value interface{} `yaml:"value" json:"value" xml:"value"`
}

func (p *Param) ToJson() (string, error) {
	d, err := utils.StructureToJson(p)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (p *Param) String() string {
	s, err := p.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (p *Param) FromJson(d string) error {
	return utils.JsonToStructure(d, p)
}

func (p *Param) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, p)
}

type Value struct {
	Name  string      `yaml:"name" json:"name" xml:"name"`
	Value interface{} `yaml:"value" json:"value" xml:"value"`
	Valid bool        `yaml:"valid" json:"valid" xml:"valid"`
}

func (v *Value) ToJson() (string, error) {
	d, err := utils.StructureToJson(v)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (v *Value) String() string {
	s, err := v.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (v *Value) FromJson(d string) error {
	return utils.JsonToStructure(d, v)
}

func (v *Value) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, v)
}

type ValueSet struct {
	File  string  `yaml:"file" json:"file" xml:"file"`
	Value []Value `yaml:"values" json:"values" xml:"value"`
}

func (v *ValueSet) ToJson() (string, error) {
	d, err := utils.StructureToJson(v)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (v *ValueSet) String() string {
	s, err := v.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (v *ValueSet) FromJson(d string) error {
	return utils.JsonToStructure(d, v)
}

func (v *ValueSet) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, v)
}

type Instance struct {
	Id         string         `yaml:"id" json:"id" xml:"id"`
	Name       string         `yaml:"name" json:"name" xml:"name"`
	Version    ProjectVersion `yaml:"version" json:"version" xml:"version"`
	State      State          `yaml:"state" json:"state" xml:"state"`
	Parameters []Param        `yaml:"parameters,omitempty" json:"parameters,omitempty" xml:"parameter,omitempty"`
	Values     ValueSet       `yaml:"values" json:"values" xml:"values"`
}

func (i *Instance) ToJson() (string, error) {
	d, err := utils.StructureToJson(i)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (i *Instance) String() string {
	s, err := i.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (i *Instance) FromJson(d string) error {
	return utils.JsonToStructure(d, i)
}

func (i *Instance) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, i)
}

func (i *Instance) GetValues() []Value {
	return i.values
}

type Job struct {
	ProjectId  string   `yaml:"projectId" json:"projectId" xml:"project-id"`
	VersionId  string   `yaml:"versionId" json:"versionId" xml:"version-id"`
	DocumentId string   `yaml:"documentId" json:"documentId" xml:"document-id"`
	IsChart    bool     `yaml:"isChart" json:"isChart" xml:"is-chart"`
	State      State    `yaml:"state" json:"state" xml:"state"`
	Instance   Instance `yaml:"instance" json:"instance" xml:"instance"`
}

func (j *Job) ToJson() (string, error) {
	d, err := utils.StructureToJson(j)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (j *Job) String() string {
	s, err := j.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (j *Job) FromJson(d string) error {
	return utils.JsonToStructure(d, j)
}

func (j *Job) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, j)
}

type Deploy struct {
	Id    string `yaml:"id" json:"id" xml:"id"`
	Name  string `yaml:"name" json:"name" xml:"name"`
	Job   []Job  `yaml:"jobs" json:"jobs" xml:"job"`
	State State  `yaml:"state" json:"state" xml:"state"`
}

func (d *Deploy) ToJson() (string, error) {
	dt, err := utils.StructureToJson(d)
	if err != nil {
		return "", err
	}
	return string(dt), nil
}

func (d *Deploy) String() string {
	s, err := d.ToJson()
	if err != nil {
		return fmt.Sprintf("<error:%s>", err.Error())
	}
	return s
}

func (dp *Deploy) FromJson(d string) error {
	return utils.JsonToStructure(d, dp)
}

func (dp *Deploy) LoadJson(path string) error {
	return utils.LoadStructureFromJsonFile(path, dp)
}

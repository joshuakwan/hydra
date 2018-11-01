package services

import (
	"context"

	"github.com/joshuakwan/hydra/codec"
	"github.com/joshuakwan/hydra/models"
	"github.com/joshuakwan/hydra/registry/rule"
)

// RuleService defines the REST interface for Rules
type RuleService interface {
	Create(action *models.Rule) error
	List() []*models.Rule
	Get(module, name string) (*models.Rule, bool)
}

// NewRuleService initializes a concrete RuleService implementation
func NewRuleService() RuleService {
	storage, err := rule.NewRuleStorage(codec.NewCodec("json"))
	if err != nil {
		return nil
	}
	return &ruleService{storage: storage}
}

type ruleService struct {
	storage *rule.Storage
}

func (r *ruleService) Create(rule *models.Rule) error {
	return r.storage.Create(context.Background(), rule)
}

func (r *ruleService) List() []*models.Rule {
	rl, err := r.storage.List(context.Background())
	if err != nil {
		return nil
	}
	return rl
}

func (r *ruleService) Get(module, name string) (*models.Rule, bool) {
	rule, err := r.storage.Get(context.Background(), module, name)
	if err != nil {
		return nil, false
	}
	return rule, true
}

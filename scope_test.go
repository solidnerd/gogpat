package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScopeApi(t *testing.T) {
	s := Scope{
		API:          true,
		ReadRegistry: false,
		ReadUser:     false,
		Sudo:         false,
	}
	ss := NewScope(url.Values{}, s)

	scopes := ss[gitlabPersonalAccessTokenScope]

	assert.Equal(t, 1, len(scopes), "They should have the expected length")

	assert.Contains(t, scopes, apiScope)
	assert.NotContains(t, scopes, readRegistryScope)
	assert.NotContains(t, scopes, readUserScope)
	assert.NotContains(t, scopes, sudoScope)
}

func TestNewScopeApiReadRegistry(t *testing.T) {
	s := Scope{
		API:          true,
		ReadRegistry: true,
		ReadUser:     false,
		Sudo:         false,
	}
	ss := NewScope(url.Values{}, s)

	scopes := ss[gitlabPersonalAccessTokenScope]

	assert.Equal(t, 2, len(scopes), "They should have the expected length")

	assert.Contains(t, scopes, apiScope)
	assert.Contains(t, scopes, readRegistryScope)
	assert.NotContains(t, scopes, readUserScope)
	assert.NotContains(t, scopes, sudoScope)
}

func TestNewScopeApiReadRegistryReadUser(t *testing.T) {
	s := Scope{
		API:          true,
		ReadRegistry: true,
		ReadUser:     true,
		Sudo:         false,
	}
	ss := NewScope(url.Values{}, s)

	scopes := ss[gitlabPersonalAccessTokenScope]

	assert.Equal(t, 3, len(scopes), "They should have the expected length")

	assert.Contains(t, scopes, apiScope)
	assert.Contains(t, scopes, readRegistryScope)
	assert.Contains(t, scopes, readUserScope)
	assert.NotContains(t, scopes, sudoScope)
}

func TestNewScopeFull(t *testing.T) {
	s := Scope{
		API:          true,
		ReadRegistry: true,
		ReadUser:     true,
		Sudo:         true,
	}
	ss := NewScope(url.Values{}, s)

	scopes := ss[gitlabPersonalAccessTokenScope]

	assert.Equal(t, 4, len(scopes), "They should have the expected length")

	assert.Contains(t, scopes, apiScope)
	assert.Contains(t, scopes, readRegistryScope)
	assert.Contains(t, scopes, readUserScope)
	assert.Contains(t, scopes, sudoScope)
}

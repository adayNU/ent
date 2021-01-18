package rule_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/facebook/ent/examples/privacytenant/ent"
	"github.com/facebook/ent/examples/privacytenant/ent/enttest"
	"github.com/facebook/ent/examples/privacytenant/viewer"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestDenyMutationsToDatasetsBelongingToOtherTenants(t *testing.T) {
	var client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	var adminCtx = viewer.NewContext(context.Background(), viewer.UserViewer{Role: viewer.Admin})

	var t1, err = client.Tenant.Create().
		SetName("Tenant 1").
		Save(adminCtx)
	assert.Nil(t, err)
	var t2 *ent.Tenant
	t2, err = client.Tenant.Create().
		SetName("Tenant 2").
		Save(adminCtx)
	assert.Nil(t, err)

	var u1 *ent.User
	u1, err = client.User.Create().
		SetName("Tenant 1 User").
		SetTenant(t1).
		Save(adminCtx)
	assert.Nil(t, err)

	var u2 *ent.User
	u2, err = client.User.Create().
		SetName("Tenant 2 User").
		SetTenant(t2).
		Save(adminCtx)
	assert.Nil(t, err)

	var user1Ctx = viewer.NewContext(context.Background(), viewer.UserViewer{
		T:    u1.QueryTenant().OnlyX(adminCtx),
		Role: viewer.View,
	})

	var user2Ctx = viewer.NewContext(context.Background(), viewer.UserViewer{
		T:    u2.QueryTenant().OnlyX(adminCtx),
		Role: viewer.View,
	})

	// Create Ops.
	_, err = client.Dataset.Create().SetTenant(t1).SetName("Dataset 1").Save(user1Ctx)
	assert.Nil(t, err)
	// Can't create dataset owned by different tenant.
	_, err = client.Dataset.Create().SetTenant(t2).SetName("Dataset 2").Save(user1Ctx)
	fmt.Println(err)
	assert.Error(t, err)
	// Can't create dataset w/o tenant edge.
	_, err = client.Dataset.Create().SetName("Dataset 2").Save(user1Ctx)
	fmt.Println(err)
	assert.Error(t, err)

	// Update Ops.
	_, err = client.Dataset.Update().SetName("Dataset 3").Save(user1Ctx)
	assert.Nil(t, err)
	// No ID set. Panics. (And test fails).
	_, err = client.Dataset.Update().SetName("Dataset 4").Save(user2Ctx)
	assert.Error(t, err)
}

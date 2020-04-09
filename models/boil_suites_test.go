// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Comments", testComments)
	t.Run("Companies", testCompanies)
	t.Run("Locations", testLocations)
	t.Run("Posts", testPosts)
	t.Run("Roles", testRoles)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Comments", testCommentsDelete)
	t.Run("Companies", testCompaniesDelete)
	t.Run("Locations", testLocationsDelete)
	t.Run("Posts", testPostsDelete)
	t.Run("Roles", testRolesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Comments", testCommentsQueryDeleteAll)
	t.Run("Companies", testCompaniesQueryDeleteAll)
	t.Run("Locations", testLocationsQueryDeleteAll)
	t.Run("Posts", testPostsQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Comments", testCommentsSliceDeleteAll)
	t.Run("Companies", testCompaniesSliceDeleteAll)
	t.Run("Locations", testLocationsSliceDeleteAll)
	t.Run("Posts", testPostsSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Comments", testCommentsExists)
	t.Run("Companies", testCompaniesExists)
	t.Run("Locations", testLocationsExists)
	t.Run("Posts", testPostsExists)
	t.Run("Roles", testRolesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Comments", testCommentsFind)
	t.Run("Companies", testCompaniesFind)
	t.Run("Locations", testLocationsFind)
	t.Run("Posts", testPostsFind)
	t.Run("Roles", testRolesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Comments", testCommentsBind)
	t.Run("Companies", testCompaniesBind)
	t.Run("Locations", testLocationsBind)
	t.Run("Posts", testPostsBind)
	t.Run("Roles", testRolesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Comments", testCommentsOne)
	t.Run("Companies", testCompaniesOne)
	t.Run("Locations", testLocationsOne)
	t.Run("Posts", testPostsOne)
	t.Run("Roles", testRolesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Comments", testCommentsAll)
	t.Run("Companies", testCompaniesAll)
	t.Run("Locations", testLocationsAll)
	t.Run("Posts", testPostsAll)
	t.Run("Roles", testRolesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Comments", testCommentsCount)
	t.Run("Companies", testCompaniesCount)
	t.Run("Locations", testLocationsCount)
	t.Run("Posts", testPostsCount)
	t.Run("Roles", testRolesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Comments", testCommentsHooks)
	t.Run("Companies", testCompaniesHooks)
	t.Run("Locations", testLocationsHooks)
	t.Run("Posts", testPostsHooks)
	t.Run("Roles", testRolesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Comments", testCommentsInsert)
	t.Run("Comments", testCommentsInsertWhitelist)
	t.Run("Companies", testCompaniesInsert)
	t.Run("Companies", testCompaniesInsertWhitelist)
	t.Run("Locations", testLocationsInsert)
	t.Run("Locations", testLocationsInsertWhitelist)
	t.Run("Posts", testPostsInsert)
	t.Run("Posts", testPostsInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CommentToPostUsingPost", testCommentToOnePostUsingPost)
	t.Run("CommentToUserUsingUser", testCommentToOneUserUsingUser)
	t.Run("LocationToCompanyUsingCompany", testLocationToOneCompanyUsingCompany)
	t.Run("PostToUserUsingUser", testPostToOneUserUsingUser)
	t.Run("UserToCompanyUsingCompany", testUserToOneCompanyUsingCompany)
	t.Run("UserToLocationUsingLocation", testUserToOneLocationUsingLocation)
	t.Run("UserToRoleUsingRole", testUserToOneRoleUsingRole)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CompanyToLocations", testCompanyToManyLocations)
	t.Run("CompanyToUsers", testCompanyToManyUsers)
	t.Run("LocationToUsers", testLocationToManyUsers)
	t.Run("PostToComments", testPostToManyComments)
	t.Run("RoleToUsers", testRoleToManyUsers)
	t.Run("UserToComments", testUserToManyComments)
	t.Run("UserToFollowerUsers", testUserToManyFollowerUsers)
	t.Run("UserToFolloweeUsers", testUserToManyFolloweeUsers)
	t.Run("UserToPosts", testUserToManyPosts)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CommentToPostUsingComments", testCommentToOneSetOpPostUsingPost)
	t.Run("CommentToUserUsingComments", testCommentToOneSetOpUserUsingUser)
	t.Run("LocationToCompanyUsingLocations", testLocationToOneSetOpCompanyUsingCompany)
	t.Run("PostToUserUsingPosts", testPostToOneSetOpUserUsingUser)
	t.Run("UserToCompanyUsingUsers", testUserToOneSetOpCompanyUsingCompany)
	t.Run("UserToLocationUsingUsers", testUserToOneSetOpLocationUsingLocation)
	t.Run("UserToRoleUsingUsers", testUserToOneSetOpRoleUsingRole)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("UserToCompanyUsingUsers", testUserToOneRemoveOpCompanyUsingCompany)
	t.Run("UserToLocationUsingUsers", testUserToOneRemoveOpLocationUsingLocation)
	t.Run("UserToRoleUsingUsers", testUserToOneRemoveOpRoleUsingRole)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CompanyToLocations", testCompanyToManyAddOpLocations)
	t.Run("CompanyToUsers", testCompanyToManyAddOpUsers)
	t.Run("LocationToUsers", testLocationToManyAddOpUsers)
	t.Run("PostToComments", testPostToManyAddOpComments)
	t.Run("RoleToUsers", testRoleToManyAddOpUsers)
	t.Run("UserToComments", testUserToManyAddOpComments)
	t.Run("UserToFollowerUsers", testUserToManyAddOpFollowerUsers)
	t.Run("UserToFolloweeUsers", testUserToManyAddOpFolloweeUsers)
	t.Run("UserToPosts", testUserToManyAddOpPosts)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("CompanyToUsers", testCompanyToManySetOpUsers)
	t.Run("LocationToUsers", testLocationToManySetOpUsers)
	t.Run("RoleToUsers", testRoleToManySetOpUsers)
	t.Run("UserToFollowerUsers", testUserToManySetOpFollowerUsers)
	t.Run("UserToFolloweeUsers", testUserToManySetOpFolloweeUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("CompanyToUsers", testCompanyToManyRemoveOpUsers)
	t.Run("LocationToUsers", testLocationToManyRemoveOpUsers)
	t.Run("RoleToUsers", testRoleToManyRemoveOpUsers)
	t.Run("UserToFollowerUsers", testUserToManyRemoveOpFollowerUsers)
	t.Run("UserToFolloweeUsers", testUserToManyRemoveOpFolloweeUsers)
}

func TestReload(t *testing.T) {
	t.Run("Comments", testCommentsReload)
	t.Run("Companies", testCompaniesReload)
	t.Run("Locations", testLocationsReload)
	t.Run("Posts", testPostsReload)
	t.Run("Roles", testRolesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Comments", testCommentsReloadAll)
	t.Run("Companies", testCompaniesReloadAll)
	t.Run("Locations", testLocationsReloadAll)
	t.Run("Posts", testPostsReloadAll)
	t.Run("Roles", testRolesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Comments", testCommentsSelect)
	t.Run("Companies", testCompaniesSelect)
	t.Run("Locations", testLocationsSelect)
	t.Run("Posts", testPostsSelect)
	t.Run("Roles", testRolesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Comments", testCommentsUpdate)
	t.Run("Companies", testCompaniesUpdate)
	t.Run("Locations", testLocationsUpdate)
	t.Run("Posts", testPostsUpdate)
	t.Run("Roles", testRolesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Comments", testCommentsSliceUpdateAll)
	t.Run("Companies", testCompaniesSliceUpdateAll)
	t.Run("Locations", testLocationsSliceUpdateAll)
	t.Run("Posts", testPostsSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}

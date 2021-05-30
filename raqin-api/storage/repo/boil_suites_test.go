// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repo

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Activities", testActivities)
	t.Run("Authors", testAuthors)
	t.Run("Books", testBooks)
	t.Run("BookAuthors", testBookAuthors)
	t.Run("BookCategories", testBookCategories)
	t.Run("BookInitiaters", testBookInitiaters)
	t.Run("Categories", testCategories)
	t.Run("Pages", testPages)
	t.Run("PageRevisions", testPageRevisions)
	t.Run("PageRevisionComments", testPageRevisionComments)
	t.Run("PageRevisionReactions", testPageRevisionReactions)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Activities", testActivitiesDelete)
	t.Run("Authors", testAuthorsDelete)
	t.Run("Books", testBooksDelete)
	t.Run("BookAuthors", testBookAuthorsDelete)
	t.Run("BookCategories", testBookCategoriesDelete)
	t.Run("BookInitiaters", testBookInitiatersDelete)
	t.Run("Categories", testCategoriesDelete)
	t.Run("Pages", testPagesDelete)
	t.Run("PageRevisions", testPageRevisionsDelete)
	t.Run("PageRevisionComments", testPageRevisionCommentsDelete)
	t.Run("PageRevisionReactions", testPageRevisionReactionsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Activities", testActivitiesQueryDeleteAll)
	t.Run("Authors", testAuthorsQueryDeleteAll)
	t.Run("Books", testBooksQueryDeleteAll)
	t.Run("BookAuthors", testBookAuthorsQueryDeleteAll)
	t.Run("BookCategories", testBookCategoriesQueryDeleteAll)
	t.Run("BookInitiaters", testBookInitiatersQueryDeleteAll)
	t.Run("Categories", testCategoriesQueryDeleteAll)
	t.Run("Pages", testPagesQueryDeleteAll)
	t.Run("PageRevisions", testPageRevisionsQueryDeleteAll)
	t.Run("PageRevisionComments", testPageRevisionCommentsQueryDeleteAll)
	t.Run("PageRevisionReactions", testPageRevisionReactionsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Activities", testActivitiesSliceDeleteAll)
	t.Run("Authors", testAuthorsSliceDeleteAll)
	t.Run("Books", testBooksSliceDeleteAll)
	t.Run("BookAuthors", testBookAuthorsSliceDeleteAll)
	t.Run("BookCategories", testBookCategoriesSliceDeleteAll)
	t.Run("BookInitiaters", testBookInitiatersSliceDeleteAll)
	t.Run("Categories", testCategoriesSliceDeleteAll)
	t.Run("Pages", testPagesSliceDeleteAll)
	t.Run("PageRevisions", testPageRevisionsSliceDeleteAll)
	t.Run("PageRevisionComments", testPageRevisionCommentsSliceDeleteAll)
	t.Run("PageRevisionReactions", testPageRevisionReactionsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Activities", testActivitiesExists)
	t.Run("Authors", testAuthorsExists)
	t.Run("Books", testBooksExists)
	t.Run("BookAuthors", testBookAuthorsExists)
	t.Run("BookCategories", testBookCategoriesExists)
	t.Run("BookInitiaters", testBookInitiatersExists)
	t.Run("Categories", testCategoriesExists)
	t.Run("Pages", testPagesExists)
	t.Run("PageRevisions", testPageRevisionsExists)
	t.Run("PageRevisionComments", testPageRevisionCommentsExists)
	t.Run("PageRevisionReactions", testPageRevisionReactionsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Activities", testActivitiesFind)
	t.Run("Authors", testAuthorsFind)
	t.Run("Books", testBooksFind)
	t.Run("BookAuthors", testBookAuthorsFind)
	t.Run("BookCategories", testBookCategoriesFind)
	t.Run("BookInitiaters", testBookInitiatersFind)
	t.Run("Categories", testCategoriesFind)
	t.Run("Pages", testPagesFind)
	t.Run("PageRevisions", testPageRevisionsFind)
	t.Run("PageRevisionComments", testPageRevisionCommentsFind)
	t.Run("PageRevisionReactions", testPageRevisionReactionsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Activities", testActivitiesBind)
	t.Run("Authors", testAuthorsBind)
	t.Run("Books", testBooksBind)
	t.Run("BookAuthors", testBookAuthorsBind)
	t.Run("BookCategories", testBookCategoriesBind)
	t.Run("BookInitiaters", testBookInitiatersBind)
	t.Run("Categories", testCategoriesBind)
	t.Run("Pages", testPagesBind)
	t.Run("PageRevisions", testPageRevisionsBind)
	t.Run("PageRevisionComments", testPageRevisionCommentsBind)
	t.Run("PageRevisionReactions", testPageRevisionReactionsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Activities", testActivitiesOne)
	t.Run("Authors", testAuthorsOne)
	t.Run("Books", testBooksOne)
	t.Run("BookAuthors", testBookAuthorsOne)
	t.Run("BookCategories", testBookCategoriesOne)
	t.Run("BookInitiaters", testBookInitiatersOne)
	t.Run("Categories", testCategoriesOne)
	t.Run("Pages", testPagesOne)
	t.Run("PageRevisions", testPageRevisionsOne)
	t.Run("PageRevisionComments", testPageRevisionCommentsOne)
	t.Run("PageRevisionReactions", testPageRevisionReactionsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Activities", testActivitiesAll)
	t.Run("Authors", testAuthorsAll)
	t.Run("Books", testBooksAll)
	t.Run("BookAuthors", testBookAuthorsAll)
	t.Run("BookCategories", testBookCategoriesAll)
	t.Run("BookInitiaters", testBookInitiatersAll)
	t.Run("Categories", testCategoriesAll)
	t.Run("Pages", testPagesAll)
	t.Run("PageRevisions", testPageRevisionsAll)
	t.Run("PageRevisionComments", testPageRevisionCommentsAll)
	t.Run("PageRevisionReactions", testPageRevisionReactionsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Activities", testActivitiesCount)
	t.Run("Authors", testAuthorsCount)
	t.Run("Books", testBooksCount)
	t.Run("BookAuthors", testBookAuthorsCount)
	t.Run("BookCategories", testBookCategoriesCount)
	t.Run("BookInitiaters", testBookInitiatersCount)
	t.Run("Categories", testCategoriesCount)
	t.Run("Pages", testPagesCount)
	t.Run("PageRevisions", testPageRevisionsCount)
	t.Run("PageRevisionComments", testPageRevisionCommentsCount)
	t.Run("PageRevisionReactions", testPageRevisionReactionsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Activities", testActivitiesHooks)
	t.Run("Authors", testAuthorsHooks)
	t.Run("Books", testBooksHooks)
	t.Run("BookAuthors", testBookAuthorsHooks)
	t.Run("BookCategories", testBookCategoriesHooks)
	t.Run("BookInitiaters", testBookInitiatersHooks)
	t.Run("Categories", testCategoriesHooks)
	t.Run("Pages", testPagesHooks)
	t.Run("PageRevisions", testPageRevisionsHooks)
	t.Run("PageRevisionComments", testPageRevisionCommentsHooks)
	t.Run("PageRevisionReactions", testPageRevisionReactionsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Activities", testActivitiesInsert)
	t.Run("Activities", testActivitiesInsertWhitelist)
	t.Run("Authors", testAuthorsInsert)
	t.Run("Authors", testAuthorsInsertWhitelist)
	t.Run("Books", testBooksInsert)
	t.Run("Books", testBooksInsertWhitelist)
	t.Run("BookAuthors", testBookAuthorsInsert)
	t.Run("BookAuthors", testBookAuthorsInsertWhitelist)
	t.Run("BookCategories", testBookCategoriesInsert)
	t.Run("BookCategories", testBookCategoriesInsertWhitelist)
	t.Run("BookInitiaters", testBookInitiatersInsert)
	t.Run("BookInitiaters", testBookInitiatersInsertWhitelist)
	t.Run("Categories", testCategoriesInsert)
	t.Run("Categories", testCategoriesInsertWhitelist)
	t.Run("Pages", testPagesInsert)
	t.Run("Pages", testPagesInsertWhitelist)
	t.Run("PageRevisions", testPageRevisionsInsert)
	t.Run("PageRevisions", testPageRevisionsInsertWhitelist)
	t.Run("PageRevisionComments", testPageRevisionCommentsInsert)
	t.Run("PageRevisionComments", testPageRevisionCommentsInsertWhitelist)
	t.Run("PageRevisionReactions", testPageRevisionReactionsInsert)
	t.Run("PageRevisionReactions", testPageRevisionReactionsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ActivityToUserUsingUser", testActivityToOneUserUsingUser)
	t.Run("BookAuthorToBookUsingBook", testBookAuthorToOneBookUsingBook)
	t.Run("BookAuthorToAuthorUsingAuthor", testBookAuthorToOneAuthorUsingAuthor)
	t.Run("BookCategoryToBookUsingBook", testBookCategoryToOneBookUsingBook)
	t.Run("BookCategoryToCategoryUsingCategory", testBookCategoryToOneCategoryUsingCategory)
	t.Run("BookInitiaterToUserUsingUser", testBookInitiaterToOneUserUsingUser)
	t.Run("BookInitiaterToBookUsingBook", testBookInitiaterToOneBookUsingBook)
	t.Run("PageToBookUsingBook", testPageToOneBookUsingBook)
	t.Run("PageToPageRevisionUsingApprovedRevisionPageRevision", testPageToOnePageRevisionUsingApprovedRevisionPageRevision)
	t.Run("PageRevisionToUserUsingReviewer", testPageRevisionToOneUserUsingReviewer)
	t.Run("PageRevisionToPageUsingPage", testPageRevisionToOnePageUsingPage)
	t.Run("PageRevisionCommentToPageRevisionUsingPageRevision", testPageRevisionCommentToOnePageRevisionUsingPageRevision)
	t.Run("PageRevisionCommentToUserUsingCommenter", testPageRevisionCommentToOneUserUsingCommenter)
	t.Run("PageRevisionReactionToPageRevisionUsingPageRevision", testPageRevisionReactionToOnePageRevisionUsingPageRevision)
	t.Run("PageRevisionReactionToUserUsingReactor", testPageRevisionReactionToOneUserUsingReactor)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AuthorToBookAuthors", testAuthorToManyBookAuthors)
	t.Run("BookToBookAuthors", testBookToManyBookAuthors)
	t.Run("BookToBookCategories", testBookToManyBookCategories)
	t.Run("BookToBookInitiaters", testBookToManyBookInitiaters)
	t.Run("BookToPages", testBookToManyPages)
	t.Run("CategoryToBookCategories", testCategoryToManyBookCategories)
	t.Run("PageToPageRevisions", testPageToManyPageRevisions)
	t.Run("PageRevisionToApprovedRevisionPages", testPageRevisionToManyApprovedRevisionPages)
	t.Run("PageRevisionToPageRevisionComments", testPageRevisionToManyPageRevisionComments)
	t.Run("PageRevisionToPageRevisionReactions", testPageRevisionToManyPageRevisionReactions)
	t.Run("UserToActivities", testUserToManyActivities)
	t.Run("UserToBookInitiaters", testUserToManyBookInitiaters)
	t.Run("UserToReviewerPageRevisions", testUserToManyReviewerPageRevisions)
	t.Run("UserToCommenterPageRevisionComments", testUserToManyCommenterPageRevisionComments)
	t.Run("UserToReactorPageRevisionReactions", testUserToManyReactorPageRevisionReactions)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ActivityToUserUsingActivities", testActivityToOneSetOpUserUsingUser)
	t.Run("BookAuthorToBookUsingBookAuthors", testBookAuthorToOneSetOpBookUsingBook)
	t.Run("BookAuthorToAuthorUsingBookAuthors", testBookAuthorToOneSetOpAuthorUsingAuthor)
	t.Run("BookCategoryToBookUsingBookCategories", testBookCategoryToOneSetOpBookUsingBook)
	t.Run("BookCategoryToCategoryUsingBookCategories", testBookCategoryToOneSetOpCategoryUsingCategory)
	t.Run("BookInitiaterToUserUsingBookInitiaters", testBookInitiaterToOneSetOpUserUsingUser)
	t.Run("BookInitiaterToBookUsingBookInitiaters", testBookInitiaterToOneSetOpBookUsingBook)
	t.Run("PageToBookUsingPages", testPageToOneSetOpBookUsingBook)
	t.Run("PageToPageRevisionUsingApprovedRevisionPages", testPageToOneSetOpPageRevisionUsingApprovedRevisionPageRevision)
	t.Run("PageRevisionToUserUsingReviewerPageRevisions", testPageRevisionToOneSetOpUserUsingReviewer)
	t.Run("PageRevisionToPageUsingPageRevisions", testPageRevisionToOneSetOpPageUsingPage)
	t.Run("PageRevisionCommentToPageRevisionUsingPageRevisionComments", testPageRevisionCommentToOneSetOpPageRevisionUsingPageRevision)
	t.Run("PageRevisionCommentToUserUsingCommenterPageRevisionComments", testPageRevisionCommentToOneSetOpUserUsingCommenter)
	t.Run("PageRevisionReactionToPageRevisionUsingPageRevisionReactions", testPageRevisionReactionToOneSetOpPageRevisionUsingPageRevision)
	t.Run("PageRevisionReactionToUserUsingReactorPageRevisionReactions", testPageRevisionReactionToOneSetOpUserUsingReactor)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("PageToPageRevisionUsingApprovedRevisionPages", testPageToOneRemoveOpPageRevisionUsingApprovedRevisionPageRevision)
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
	t.Run("AuthorToBookAuthors", testAuthorToManyAddOpBookAuthors)
	t.Run("BookToBookAuthors", testBookToManyAddOpBookAuthors)
	t.Run("BookToBookCategories", testBookToManyAddOpBookCategories)
	t.Run("BookToBookInitiaters", testBookToManyAddOpBookInitiaters)
	t.Run("BookToPages", testBookToManyAddOpPages)
	t.Run("CategoryToBookCategories", testCategoryToManyAddOpBookCategories)
	t.Run("PageToPageRevisions", testPageToManyAddOpPageRevisions)
	t.Run("PageRevisionToApprovedRevisionPages", testPageRevisionToManyAddOpApprovedRevisionPages)
	t.Run("PageRevisionToPageRevisionComments", testPageRevisionToManyAddOpPageRevisionComments)
	t.Run("PageRevisionToPageRevisionReactions", testPageRevisionToManyAddOpPageRevisionReactions)
	t.Run("UserToActivities", testUserToManyAddOpActivities)
	t.Run("UserToBookInitiaters", testUserToManyAddOpBookInitiaters)
	t.Run("UserToReviewerPageRevisions", testUserToManyAddOpReviewerPageRevisions)
	t.Run("UserToCommenterPageRevisionComments", testUserToManyAddOpCommenterPageRevisionComments)
	t.Run("UserToReactorPageRevisionReactions", testUserToManyAddOpReactorPageRevisionReactions)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("PageRevisionToApprovedRevisionPages", testPageRevisionToManySetOpApprovedRevisionPages)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("PageRevisionToApprovedRevisionPages", testPageRevisionToManyRemoveOpApprovedRevisionPages)
}

func TestReload(t *testing.T) {
	t.Run("Activities", testActivitiesReload)
	t.Run("Authors", testAuthorsReload)
	t.Run("Books", testBooksReload)
	t.Run("BookAuthors", testBookAuthorsReload)
	t.Run("BookCategories", testBookCategoriesReload)
	t.Run("BookInitiaters", testBookInitiatersReload)
	t.Run("Categories", testCategoriesReload)
	t.Run("Pages", testPagesReload)
	t.Run("PageRevisions", testPageRevisionsReload)
	t.Run("PageRevisionComments", testPageRevisionCommentsReload)
	t.Run("PageRevisionReactions", testPageRevisionReactionsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Activities", testActivitiesReloadAll)
	t.Run("Authors", testAuthorsReloadAll)
	t.Run("Books", testBooksReloadAll)
	t.Run("BookAuthors", testBookAuthorsReloadAll)
	t.Run("BookCategories", testBookCategoriesReloadAll)
	t.Run("BookInitiaters", testBookInitiatersReloadAll)
	t.Run("Categories", testCategoriesReloadAll)
	t.Run("Pages", testPagesReloadAll)
	t.Run("PageRevisions", testPageRevisionsReloadAll)
	t.Run("PageRevisionComments", testPageRevisionCommentsReloadAll)
	t.Run("PageRevisionReactions", testPageRevisionReactionsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Activities", testActivitiesSelect)
	t.Run("Authors", testAuthorsSelect)
	t.Run("Books", testBooksSelect)
	t.Run("BookAuthors", testBookAuthorsSelect)
	t.Run("BookCategories", testBookCategoriesSelect)
	t.Run("BookInitiaters", testBookInitiatersSelect)
	t.Run("Categories", testCategoriesSelect)
	t.Run("Pages", testPagesSelect)
	t.Run("PageRevisions", testPageRevisionsSelect)
	t.Run("PageRevisionComments", testPageRevisionCommentsSelect)
	t.Run("PageRevisionReactions", testPageRevisionReactionsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Activities", testActivitiesUpdate)
	t.Run("Authors", testAuthorsUpdate)
	t.Run("Books", testBooksUpdate)
	t.Run("BookAuthors", testBookAuthorsUpdate)
	t.Run("BookCategories", testBookCategoriesUpdate)
	t.Run("BookInitiaters", testBookInitiatersUpdate)
	t.Run("Categories", testCategoriesUpdate)
	t.Run("Pages", testPagesUpdate)
	t.Run("PageRevisions", testPageRevisionsUpdate)
	t.Run("PageRevisionComments", testPageRevisionCommentsUpdate)
	t.Run("PageRevisionReactions", testPageRevisionReactionsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Activities", testActivitiesSliceUpdateAll)
	t.Run("Authors", testAuthorsSliceUpdateAll)
	t.Run("Books", testBooksSliceUpdateAll)
	t.Run("BookAuthors", testBookAuthorsSliceUpdateAll)
	t.Run("BookCategories", testBookCategoriesSliceUpdateAll)
	t.Run("BookInitiaters", testBookInitiatersSliceUpdateAll)
	t.Run("Categories", testCategoriesSliceUpdateAll)
	t.Run("Pages", testPagesSliceUpdateAll)
	t.Run("PageRevisions", testPageRevisionsSliceUpdateAll)
	t.Run("PageRevisionComments", testPageRevisionCommentsSliceUpdateAll)
	t.Run("PageRevisionReactions", testPageRevisionReactionsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}

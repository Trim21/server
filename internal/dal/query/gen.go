// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

func Use(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		AccessToken:       newAccessToken(db),
		App:               newApp(db),
		Cast:              newCast(db),
		Character:         newCharacter(db),
		CharacterSubjects: newCharacterSubjects(db),
		Episode:           newEpisode(db),
		Group:             newGroup(db),
		GroupMember:       newGroupMember(db),
		Index:             newIndex(db),
		IndexSubject:      newIndexSubject(db),
		Member:            newMember(db),
		OAuthClient:       newOAuthClient(db),
		Person:            newPerson(db),
		PersonField:       newPersonField(db),
		PersonSubjects:    newPersonSubjects(db),
		RevisionHistory:   newRevisionHistory(db),
		RevisionText:      newRevisionText(db),
		Subject:           newSubject(db),
		SubjectCollection: newSubjectCollection(db),
		SubjectField:      newSubjectField(db),
		SubjectRelation:   newSubjectRelation(db),
		SubjectRevision:   newSubjectRevision(db),
		UserGroup:         newUserGroup(db),
		WebSession:        newWebSession(db),
	}
}

type Query struct {
	db *gorm.DB

	AccessToken       accessToken
	App               app
	Cast              cast
	Character         character
	CharacterSubjects characterSubjects
	Episode           episode
	Group             group
	GroupMember       groupMember
	Index             index
	IndexSubject      indexSubject
	Member            member
	OAuthClient       oAuthClient
	Person            person
	PersonField       personField
	PersonSubjects    personSubjects
	RevisionHistory   revisionHistory
	RevisionText      revisionText
	Subject           subject
	SubjectCollection subjectCollection
	SubjectField      subjectField
	SubjectRelation   subjectRelation
	SubjectRevision   subjectRevision
	UserGroup         userGroup
	WebSession        webSession
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		AccessToken:       q.AccessToken.clone(db),
		App:               q.App.clone(db),
		Cast:              q.Cast.clone(db),
		Character:         q.Character.clone(db),
		CharacterSubjects: q.CharacterSubjects.clone(db),
		Episode:           q.Episode.clone(db),
		Group:             q.Group.clone(db),
		GroupMember:       q.GroupMember.clone(db),
		Index:             q.Index.clone(db),
		IndexSubject:      q.IndexSubject.clone(db),
		Member:            q.Member.clone(db),
		OAuthClient:       q.OAuthClient.clone(db),
		Person:            q.Person.clone(db),
		PersonField:       q.PersonField.clone(db),
		PersonSubjects:    q.PersonSubjects.clone(db),
		RevisionHistory:   q.RevisionHistory.clone(db),
		RevisionText:      q.RevisionText.clone(db),
		Subject:           q.Subject.clone(db),
		SubjectCollection: q.SubjectCollection.clone(db),
		SubjectField:      q.SubjectField.clone(db),
		SubjectRelation:   q.SubjectRelation.clone(db),
		SubjectRevision:   q.SubjectRevision.clone(db),
		UserGroup:         q.UserGroup.clone(db),
		WebSession:        q.WebSession.clone(db),
	}
}

type queryCtx struct {
	AccessToken       *accessTokenDo
	App               *appDo
	Cast              *castDo
	Character         *characterDo
	CharacterSubjects *characterSubjectsDo
	Episode           *episodeDo
	Group             *groupDo
	GroupMember       *groupMemberDo
	Index             *indexDo
	IndexSubject      *indexSubjectDo
	Member            *memberDo
	OAuthClient       *oAuthClientDo
	Person            *personDo
	PersonField       *personFieldDo
	PersonSubjects    *personSubjectsDo
	RevisionHistory   *revisionHistoryDo
	RevisionText      *revisionTextDo
	Subject           *subjectDo
	SubjectCollection *subjectCollectionDo
	SubjectField      *subjectFieldDo
	SubjectRelation   *subjectRelationDo
	SubjectRevision   *subjectRevisionDo
	UserGroup         *userGroupDo
	WebSession        *webSessionDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AccessToken:       q.AccessToken.WithContext(ctx),
		App:               q.App.WithContext(ctx),
		Cast:              q.Cast.WithContext(ctx),
		Character:         q.Character.WithContext(ctx),
		CharacterSubjects: q.CharacterSubjects.WithContext(ctx),
		Episode:           q.Episode.WithContext(ctx),
		Group:             q.Group.WithContext(ctx),
		GroupMember:       q.GroupMember.WithContext(ctx),
		Index:             q.Index.WithContext(ctx),
		IndexSubject:      q.IndexSubject.WithContext(ctx),
		Member:            q.Member.WithContext(ctx),
		OAuthClient:       q.OAuthClient.WithContext(ctx),
		Person:            q.Person.WithContext(ctx),
		PersonField:       q.PersonField.WithContext(ctx),
		PersonSubjects:    q.PersonSubjects.WithContext(ctx),
		RevisionHistory:   q.RevisionHistory.WithContext(ctx),
		RevisionText:      q.RevisionText.WithContext(ctx),
		Subject:           q.Subject.WithContext(ctx),
		SubjectCollection: q.SubjectCollection.WithContext(ctx),
		SubjectField:      q.SubjectField.WithContext(ctx),
		SubjectRelation:   q.SubjectRelation.WithContext(ctx),
		SubjectRevision:   q.SubjectRevision.WithContext(ctx),
		UserGroup:         q.UserGroup.WithContext(ctx),
		WebSession:        q.WebSession.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}

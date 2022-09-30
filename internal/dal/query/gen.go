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
		db:                  db,
		AccessToken:         newAccessToken(db),
		App:                 newApp(db),
		Cast:                newCast(db),
		Character:           newCharacter(db),
		CharacterComment:    newCharacterComment(db),
		CharacterSubjects:   newCharacterSubjects(db),
		EpCollection:        newEpCollection(db),
		EpRevision:          newEpRevision(db),
		Episode:             newEpisode(db),
		EpisodeComment:      newEpisodeComment(db),
		Friend:              newFriend(db),
		Group:               newGroup(db),
		GroupMember:         newGroupMember(db),
		GroupTopic:          newGroupTopic(db),
		GroupTopicComment:   newGroupTopicComment(db),
		Index:               newIndex(db),
		IndexComment:        newIndexComment(db),
		IndexSubject:        newIndexSubject(db),
		Member:              newMember(db),
		OAuthClient:         newOAuthClient(db),
		Person:              newPerson(db),
		PersonComment:       newPersonComment(db),
		PersonField:         newPersonField(db),
		PersonSubjects:      newPersonSubjects(db),
		RevisionHistory:     newRevisionHistory(db),
		RevisionText:        newRevisionText(db),
		Subject:             newSubject(db),
		SubjectCollection:   newSubjectCollection(db),
		SubjectField:        newSubjectField(db),
		SubjectRelation:     newSubjectRelation(db),
		SubjectRevision:     newSubjectRevision(db),
		SubjectTopic:        newSubjectTopic(db),
		SubjectTopicComment: newSubjectTopicComment(db),
		TimeLine:            newTimeLine(db),
		UserGroup:           newUserGroup(db),
		WebSession:          newWebSession(db),
	}
}

type Query struct {
	db *gorm.DB

	AccessToken         accessToken
	App                 app
	Cast                cast
	Character           character
	CharacterComment    characterComment
	CharacterSubjects   characterSubjects
	EpCollection        epCollection
	EpRevision          epRevision
	Episode             episode
	EpisodeComment      episodeComment
	Friend              friend
	Group               group
	GroupMember         groupMember
	GroupTopic          groupTopic
	GroupTopicComment   groupTopicComment
	Index               index
	IndexComment        indexComment
	IndexSubject        indexSubject
	Member              member
	OAuthClient         oAuthClient
	Person              person
	PersonComment       personComment
	PersonField         personField
	PersonSubjects      personSubjects
	RevisionHistory     revisionHistory
	RevisionText        revisionText
	Subject             subject
	SubjectCollection   subjectCollection
	SubjectField        subjectField
	SubjectRelation     subjectRelation
	SubjectRevision     subjectRevision
	SubjectTopic        subjectTopic
	SubjectTopicComment subjectTopicComment
	TimeLine            timeLine
	UserGroup           userGroup
	WebSession          webSession
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                  db,
		AccessToken:         q.AccessToken.clone(db),
		App:                 q.App.clone(db),
		Cast:                q.Cast.clone(db),
		Character:           q.Character.clone(db),
		CharacterComment:    q.CharacterComment.clone(db),
		CharacterSubjects:   q.CharacterSubjects.clone(db),
		EpCollection:        q.EpCollection.clone(db),
		EpRevision:          q.EpRevision.clone(db),
		Episode:             q.Episode.clone(db),
		EpisodeComment:      q.EpisodeComment.clone(db),
		Friend:              q.Friend.clone(db),
		Group:               q.Group.clone(db),
		GroupMember:         q.GroupMember.clone(db),
		GroupTopic:          q.GroupTopic.clone(db),
		GroupTopicComment:   q.GroupTopicComment.clone(db),
		Index:               q.Index.clone(db),
		IndexComment:        q.IndexComment.clone(db),
		IndexSubject:        q.IndexSubject.clone(db),
		Member:              q.Member.clone(db),
		OAuthClient:         q.OAuthClient.clone(db),
		Person:              q.Person.clone(db),
		PersonComment:       q.PersonComment.clone(db),
		PersonField:         q.PersonField.clone(db),
		PersonSubjects:      q.PersonSubjects.clone(db),
		RevisionHistory:     q.RevisionHistory.clone(db),
		RevisionText:        q.RevisionText.clone(db),
		Subject:             q.Subject.clone(db),
		SubjectCollection:   q.SubjectCollection.clone(db),
		SubjectField:        q.SubjectField.clone(db),
		SubjectRelation:     q.SubjectRelation.clone(db),
		SubjectRevision:     q.SubjectRevision.clone(db),
		SubjectTopic:        q.SubjectTopic.clone(db),
		SubjectTopicComment: q.SubjectTopicComment.clone(db),
		TimeLine:            q.TimeLine.clone(db),
		UserGroup:           q.UserGroup.clone(db),
		WebSession:          q.WebSession.clone(db),
	}
}

type queryCtx struct {
	AccessToken         *accessTokenDo
	App                 *appDo
	Cast                *castDo
	Character           *characterDo
	CharacterComment    *characterCommentDo
	CharacterSubjects   *characterSubjectsDo
	EpCollection        *epCollectionDo
	EpRevision          *epRevisionDo
	Episode             *episodeDo
	EpisodeComment      *episodeCommentDo
	Friend              *friendDo
	Group               *groupDo
	GroupMember         *groupMemberDo
	GroupTopic          *groupTopicDo
	GroupTopicComment   *groupTopicCommentDo
	Index               *indexDo
	IndexComment        *indexCommentDo
	IndexSubject        *indexSubjectDo
	Member              *memberDo
	OAuthClient         *oAuthClientDo
	Person              *personDo
	PersonComment       *personCommentDo
	PersonField         *personFieldDo
	PersonSubjects      *personSubjectsDo
	RevisionHistory     *revisionHistoryDo
	RevisionText        *revisionTextDo
	Subject             *subjectDo
	SubjectCollection   *subjectCollectionDo
	SubjectField        *subjectFieldDo
	SubjectRelation     *subjectRelationDo
	SubjectRevision     *subjectRevisionDo
	SubjectTopic        *subjectTopicDo
	SubjectTopicComment *subjectTopicCommentDo
	TimeLine            *timeLineDo
	UserGroup           *userGroupDo
	WebSession          *webSessionDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AccessToken:         q.AccessToken.WithContext(ctx),
		App:                 q.App.WithContext(ctx),
		Cast:                q.Cast.WithContext(ctx),
		Character:           q.Character.WithContext(ctx),
		CharacterComment:    q.CharacterComment.WithContext(ctx),
		CharacterSubjects:   q.CharacterSubjects.WithContext(ctx),
		EpCollection:        q.EpCollection.WithContext(ctx),
		EpRevision:          q.EpRevision.WithContext(ctx),
		Episode:             q.Episode.WithContext(ctx),
		EpisodeComment:      q.EpisodeComment.WithContext(ctx),
		Friend:              q.Friend.WithContext(ctx),
		Group:               q.Group.WithContext(ctx),
		GroupMember:         q.GroupMember.WithContext(ctx),
		GroupTopic:          q.GroupTopic.WithContext(ctx),
		GroupTopicComment:   q.GroupTopicComment.WithContext(ctx),
		Index:               q.Index.WithContext(ctx),
		IndexComment:        q.IndexComment.WithContext(ctx),
		IndexSubject:        q.IndexSubject.WithContext(ctx),
		Member:              q.Member.WithContext(ctx),
		OAuthClient:         q.OAuthClient.WithContext(ctx),
		Person:              q.Person.WithContext(ctx),
		PersonComment:       q.PersonComment.WithContext(ctx),
		PersonField:         q.PersonField.WithContext(ctx),
		PersonSubjects:      q.PersonSubjects.WithContext(ctx),
		RevisionHistory:     q.RevisionHistory.WithContext(ctx),
		RevisionText:        q.RevisionText.WithContext(ctx),
		Subject:             q.Subject.WithContext(ctx),
		SubjectCollection:   q.SubjectCollection.WithContext(ctx),
		SubjectField:        q.SubjectField.WithContext(ctx),
		SubjectRelation:     q.SubjectRelation.WithContext(ctx),
		SubjectRevision:     q.SubjectRevision.WithContext(ctx),
		SubjectTopic:        q.SubjectTopic.WithContext(ctx),
		SubjectTopicComment: q.SubjectTopicComment.WithContext(ctx),
		TimeLine:            q.TimeLine.WithContext(ctx),
		UserGroup:           q.UserGroup.WithContext(ctx),
		WebSession:          q.WebSession.WithContext(ctx),
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

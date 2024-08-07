// Code generated by ent, DO NOT EDIT.

package track

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Pineapple217/Sortify/web/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldName, v))
}

// Artist applies equality check predicate on the "artist" field. It's identical to ArtistEQ.
func Artist(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldArtist, v))
}

// ImgSmallURL applies equality check predicate on the "img_small_url" field. It's identical to ImgSmallURLEQ.
func ImgSmallURL(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldImgSmallURL, v))
}

// ImgMediumURL applies equality check predicate on the "img_medium_url" field. It's identical to ImgMediumURLEQ.
func ImgMediumURL(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldImgMediumURL, v))
}

// ImgLargeURL applies equality check predicate on the "img_large_url" field. It's identical to ImgLargeURLEQ.
func ImgLargeURL(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldImgLargeURL, v))
}

// ReleaseDate applies equality check predicate on the "release_date" field. It's identical to ReleaseDateEQ.
func ReleaseDate(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldReleaseDate, v))
}

// SpotifyID applies equality check predicate on the "spotify_id" field. It's identical to SpotifyIDEQ.
func SpotifyID(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldSpotifyID, v))
}

// DurationMs applies equality check predicate on the "duration_ms" field. It's identical to DurationMsEQ.
func DurationMs(v int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldDurationMs, v))
}

// PreviewURL applies equality check predicate on the "preview_url" field. It's identical to PreviewURLEQ.
func PreviewURL(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldPreviewURL, v))
}

// Popularity applies equality check predicate on the "popularity" field. It's identical to PopularityEQ.
func Popularity(v int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldPopularity, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldDeletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldName, v))
}

// ArtistEQ applies the EQ predicate on the "artist" field.
func ArtistEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldArtist, v))
}

// ArtistNEQ applies the NEQ predicate on the "artist" field.
func ArtistNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldArtist, v))
}

// ArtistIn applies the In predicate on the "artist" field.
func ArtistIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldArtist, vs...))
}

// ArtistNotIn applies the NotIn predicate on the "artist" field.
func ArtistNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldArtist, vs...))
}

// ArtistGT applies the GT predicate on the "artist" field.
func ArtistGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldArtist, v))
}

// ArtistGTE applies the GTE predicate on the "artist" field.
func ArtistGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldArtist, v))
}

// ArtistLT applies the LT predicate on the "artist" field.
func ArtistLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldArtist, v))
}

// ArtistLTE applies the LTE predicate on the "artist" field.
func ArtistLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldArtist, v))
}

// ArtistContains applies the Contains predicate on the "artist" field.
func ArtistContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldArtist, v))
}

// ArtistHasPrefix applies the HasPrefix predicate on the "artist" field.
func ArtistHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldArtist, v))
}

// ArtistHasSuffix applies the HasSuffix predicate on the "artist" field.
func ArtistHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldArtist, v))
}

// ArtistEqualFold applies the EqualFold predicate on the "artist" field.
func ArtistEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldArtist, v))
}

// ArtistContainsFold applies the ContainsFold predicate on the "artist" field.
func ArtistContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldArtist, v))
}

// ImgSmallURLEQ applies the EQ predicate on the "img_small_url" field.
func ImgSmallURLEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldImgSmallURL, v))
}

// ImgSmallURLNEQ applies the NEQ predicate on the "img_small_url" field.
func ImgSmallURLNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldImgSmallURL, v))
}

// ImgSmallURLIn applies the In predicate on the "img_small_url" field.
func ImgSmallURLIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldImgSmallURL, vs...))
}

// ImgSmallURLNotIn applies the NotIn predicate on the "img_small_url" field.
func ImgSmallURLNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldImgSmallURL, vs...))
}

// ImgSmallURLGT applies the GT predicate on the "img_small_url" field.
func ImgSmallURLGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldImgSmallURL, v))
}

// ImgSmallURLGTE applies the GTE predicate on the "img_small_url" field.
func ImgSmallURLGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldImgSmallURL, v))
}

// ImgSmallURLLT applies the LT predicate on the "img_small_url" field.
func ImgSmallURLLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldImgSmallURL, v))
}

// ImgSmallURLLTE applies the LTE predicate on the "img_small_url" field.
func ImgSmallURLLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldImgSmallURL, v))
}

// ImgSmallURLContains applies the Contains predicate on the "img_small_url" field.
func ImgSmallURLContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldImgSmallURL, v))
}

// ImgSmallURLHasPrefix applies the HasPrefix predicate on the "img_small_url" field.
func ImgSmallURLHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldImgSmallURL, v))
}

// ImgSmallURLHasSuffix applies the HasSuffix predicate on the "img_small_url" field.
func ImgSmallURLHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldImgSmallURL, v))
}

// ImgSmallURLIsNil applies the IsNil predicate on the "img_small_url" field.
func ImgSmallURLIsNil() predicate.Track {
	return predicate.Track(sql.FieldIsNull(FieldImgSmallURL))
}

// ImgSmallURLNotNil applies the NotNil predicate on the "img_small_url" field.
func ImgSmallURLNotNil() predicate.Track {
	return predicate.Track(sql.FieldNotNull(FieldImgSmallURL))
}

// ImgSmallURLEqualFold applies the EqualFold predicate on the "img_small_url" field.
func ImgSmallURLEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldImgSmallURL, v))
}

// ImgSmallURLContainsFold applies the ContainsFold predicate on the "img_small_url" field.
func ImgSmallURLContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldImgSmallURL, v))
}

// ImgMediumURLEQ applies the EQ predicate on the "img_medium_url" field.
func ImgMediumURLEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldImgMediumURL, v))
}

// ImgMediumURLNEQ applies the NEQ predicate on the "img_medium_url" field.
func ImgMediumURLNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldImgMediumURL, v))
}

// ImgMediumURLIn applies the In predicate on the "img_medium_url" field.
func ImgMediumURLIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldImgMediumURL, vs...))
}

// ImgMediumURLNotIn applies the NotIn predicate on the "img_medium_url" field.
func ImgMediumURLNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldImgMediumURL, vs...))
}

// ImgMediumURLGT applies the GT predicate on the "img_medium_url" field.
func ImgMediumURLGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldImgMediumURL, v))
}

// ImgMediumURLGTE applies the GTE predicate on the "img_medium_url" field.
func ImgMediumURLGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldImgMediumURL, v))
}

// ImgMediumURLLT applies the LT predicate on the "img_medium_url" field.
func ImgMediumURLLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldImgMediumURL, v))
}

// ImgMediumURLLTE applies the LTE predicate on the "img_medium_url" field.
func ImgMediumURLLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldImgMediumURL, v))
}

// ImgMediumURLContains applies the Contains predicate on the "img_medium_url" field.
func ImgMediumURLContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldImgMediumURL, v))
}

// ImgMediumURLHasPrefix applies the HasPrefix predicate on the "img_medium_url" field.
func ImgMediumURLHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldImgMediumURL, v))
}

// ImgMediumURLHasSuffix applies the HasSuffix predicate on the "img_medium_url" field.
func ImgMediumURLHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldImgMediumURL, v))
}

// ImgMediumURLIsNil applies the IsNil predicate on the "img_medium_url" field.
func ImgMediumURLIsNil() predicate.Track {
	return predicate.Track(sql.FieldIsNull(FieldImgMediumURL))
}

// ImgMediumURLNotNil applies the NotNil predicate on the "img_medium_url" field.
func ImgMediumURLNotNil() predicate.Track {
	return predicate.Track(sql.FieldNotNull(FieldImgMediumURL))
}

// ImgMediumURLEqualFold applies the EqualFold predicate on the "img_medium_url" field.
func ImgMediumURLEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldImgMediumURL, v))
}

// ImgMediumURLContainsFold applies the ContainsFold predicate on the "img_medium_url" field.
func ImgMediumURLContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldImgMediumURL, v))
}

// ImgLargeURLEQ applies the EQ predicate on the "img_large_url" field.
func ImgLargeURLEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldImgLargeURL, v))
}

// ImgLargeURLNEQ applies the NEQ predicate on the "img_large_url" field.
func ImgLargeURLNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldImgLargeURL, v))
}

// ImgLargeURLIn applies the In predicate on the "img_large_url" field.
func ImgLargeURLIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldImgLargeURL, vs...))
}

// ImgLargeURLNotIn applies the NotIn predicate on the "img_large_url" field.
func ImgLargeURLNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldImgLargeURL, vs...))
}

// ImgLargeURLGT applies the GT predicate on the "img_large_url" field.
func ImgLargeURLGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldImgLargeURL, v))
}

// ImgLargeURLGTE applies the GTE predicate on the "img_large_url" field.
func ImgLargeURLGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldImgLargeURL, v))
}

// ImgLargeURLLT applies the LT predicate on the "img_large_url" field.
func ImgLargeURLLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldImgLargeURL, v))
}

// ImgLargeURLLTE applies the LTE predicate on the "img_large_url" field.
func ImgLargeURLLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldImgLargeURL, v))
}

// ImgLargeURLContains applies the Contains predicate on the "img_large_url" field.
func ImgLargeURLContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldImgLargeURL, v))
}

// ImgLargeURLHasPrefix applies the HasPrefix predicate on the "img_large_url" field.
func ImgLargeURLHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldImgLargeURL, v))
}

// ImgLargeURLHasSuffix applies the HasSuffix predicate on the "img_large_url" field.
func ImgLargeURLHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldImgLargeURL, v))
}

// ImgLargeURLIsNil applies the IsNil predicate on the "img_large_url" field.
func ImgLargeURLIsNil() predicate.Track {
	return predicate.Track(sql.FieldIsNull(FieldImgLargeURL))
}

// ImgLargeURLNotNil applies the NotNil predicate on the "img_large_url" field.
func ImgLargeURLNotNil() predicate.Track {
	return predicate.Track(sql.FieldNotNull(FieldImgLargeURL))
}

// ImgLargeURLEqualFold applies the EqualFold predicate on the "img_large_url" field.
func ImgLargeURLEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldImgLargeURL, v))
}

// ImgLargeURLContainsFold applies the ContainsFold predicate on the "img_large_url" field.
func ImgLargeURLContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldImgLargeURL, v))
}

// ReleaseDateEQ applies the EQ predicate on the "release_date" field.
func ReleaseDateEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldReleaseDate, v))
}

// ReleaseDateNEQ applies the NEQ predicate on the "release_date" field.
func ReleaseDateNEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldReleaseDate, v))
}

// ReleaseDateIn applies the In predicate on the "release_date" field.
func ReleaseDateIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldReleaseDate, vs...))
}

// ReleaseDateNotIn applies the NotIn predicate on the "release_date" field.
func ReleaseDateNotIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldReleaseDate, vs...))
}

// ReleaseDateGT applies the GT predicate on the "release_date" field.
func ReleaseDateGT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldReleaseDate, v))
}

// ReleaseDateGTE applies the GTE predicate on the "release_date" field.
func ReleaseDateGTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldReleaseDate, v))
}

// ReleaseDateLT applies the LT predicate on the "release_date" field.
func ReleaseDateLT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldReleaseDate, v))
}

// ReleaseDateLTE applies the LTE predicate on the "release_date" field.
func ReleaseDateLTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldReleaseDate, v))
}

// SpotifyIDEQ applies the EQ predicate on the "spotify_id" field.
func SpotifyIDEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldSpotifyID, v))
}

// SpotifyIDNEQ applies the NEQ predicate on the "spotify_id" field.
func SpotifyIDNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldSpotifyID, v))
}

// SpotifyIDIn applies the In predicate on the "spotify_id" field.
func SpotifyIDIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldSpotifyID, vs...))
}

// SpotifyIDNotIn applies the NotIn predicate on the "spotify_id" field.
func SpotifyIDNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldSpotifyID, vs...))
}

// SpotifyIDGT applies the GT predicate on the "spotify_id" field.
func SpotifyIDGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldSpotifyID, v))
}

// SpotifyIDGTE applies the GTE predicate on the "spotify_id" field.
func SpotifyIDGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldSpotifyID, v))
}

// SpotifyIDLT applies the LT predicate on the "spotify_id" field.
func SpotifyIDLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldSpotifyID, v))
}

// SpotifyIDLTE applies the LTE predicate on the "spotify_id" field.
func SpotifyIDLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldSpotifyID, v))
}

// SpotifyIDContains applies the Contains predicate on the "spotify_id" field.
func SpotifyIDContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldSpotifyID, v))
}

// SpotifyIDHasPrefix applies the HasPrefix predicate on the "spotify_id" field.
func SpotifyIDHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldSpotifyID, v))
}

// SpotifyIDHasSuffix applies the HasSuffix predicate on the "spotify_id" field.
func SpotifyIDHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldSpotifyID, v))
}

// SpotifyIDEqualFold applies the EqualFold predicate on the "spotify_id" field.
func SpotifyIDEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldSpotifyID, v))
}

// SpotifyIDContainsFold applies the ContainsFold predicate on the "spotify_id" field.
func SpotifyIDContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldSpotifyID, v))
}

// DurationMsEQ applies the EQ predicate on the "duration_ms" field.
func DurationMsEQ(v int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldDurationMs, v))
}

// DurationMsNEQ applies the NEQ predicate on the "duration_ms" field.
func DurationMsNEQ(v int) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldDurationMs, v))
}

// DurationMsIn applies the In predicate on the "duration_ms" field.
func DurationMsIn(vs ...int) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldDurationMs, vs...))
}

// DurationMsNotIn applies the NotIn predicate on the "duration_ms" field.
func DurationMsNotIn(vs ...int) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldDurationMs, vs...))
}

// DurationMsGT applies the GT predicate on the "duration_ms" field.
func DurationMsGT(v int) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldDurationMs, v))
}

// DurationMsGTE applies the GTE predicate on the "duration_ms" field.
func DurationMsGTE(v int) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldDurationMs, v))
}

// DurationMsLT applies the LT predicate on the "duration_ms" field.
func DurationMsLT(v int) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldDurationMs, v))
}

// DurationMsLTE applies the LTE predicate on the "duration_ms" field.
func DurationMsLTE(v int) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldDurationMs, v))
}

// PreviewURLEQ applies the EQ predicate on the "preview_url" field.
func PreviewURLEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldPreviewURL, v))
}

// PreviewURLNEQ applies the NEQ predicate on the "preview_url" field.
func PreviewURLNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldPreviewURL, v))
}

// PreviewURLIn applies the In predicate on the "preview_url" field.
func PreviewURLIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldPreviewURL, vs...))
}

// PreviewURLNotIn applies the NotIn predicate on the "preview_url" field.
func PreviewURLNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldPreviewURL, vs...))
}

// PreviewURLGT applies the GT predicate on the "preview_url" field.
func PreviewURLGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldPreviewURL, v))
}

// PreviewURLGTE applies the GTE predicate on the "preview_url" field.
func PreviewURLGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldPreviewURL, v))
}

// PreviewURLLT applies the LT predicate on the "preview_url" field.
func PreviewURLLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldPreviewURL, v))
}

// PreviewURLLTE applies the LTE predicate on the "preview_url" field.
func PreviewURLLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldPreviewURL, v))
}

// PreviewURLContains applies the Contains predicate on the "preview_url" field.
func PreviewURLContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldPreviewURL, v))
}

// PreviewURLHasPrefix applies the HasPrefix predicate on the "preview_url" field.
func PreviewURLHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldPreviewURL, v))
}

// PreviewURLHasSuffix applies the HasSuffix predicate on the "preview_url" field.
func PreviewURLHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldPreviewURL, v))
}

// PreviewURLIsNil applies the IsNil predicate on the "preview_url" field.
func PreviewURLIsNil() predicate.Track {
	return predicate.Track(sql.FieldIsNull(FieldPreviewURL))
}

// PreviewURLNotNil applies the NotNil predicate on the "preview_url" field.
func PreviewURLNotNil() predicate.Track {
	return predicate.Track(sql.FieldNotNull(FieldPreviewURL))
}

// PreviewURLEqualFold applies the EqualFold predicate on the "preview_url" field.
func PreviewURLEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldPreviewURL, v))
}

// PreviewURLContainsFold applies the ContainsFold predicate on the "preview_url" field.
func PreviewURLContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldPreviewURL, v))
}

// PopularityEQ applies the EQ predicate on the "popularity" field.
func PopularityEQ(v int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldPopularity, v))
}

// PopularityNEQ applies the NEQ predicate on the "popularity" field.
func PopularityNEQ(v int) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldPopularity, v))
}

// PopularityIn applies the In predicate on the "popularity" field.
func PopularityIn(vs ...int) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldPopularity, vs...))
}

// PopularityNotIn applies the NotIn predicate on the "popularity" field.
func PopularityNotIn(vs ...int) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldPopularity, vs...))
}

// PopularityGT applies the GT predicate on the "popularity" field.
func PopularityGT(v int) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldPopularity, v))
}

// PopularityGTE applies the GTE predicate on the "popularity" field.
func PopularityGTE(v int) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldPopularity, v))
}

// PopularityLT applies the LT predicate on the "popularity" field.
func PopularityLT(v int) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldPopularity, v))
}

// PopularityLTE applies the LTE predicate on the "popularity" field.
func PopularityLTE(v int) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldPopularity, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Track {
	return predicate.Track(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Track {
	return predicate.Track(sql.FieldNotNull(FieldDeletedAt))
}

// HasPlaylists applies the HasEdge predicate on the "playlists" edge.
func HasPlaylists() predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PlaylistsTable, PlaylistsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaylistsWith applies the HasEdge predicate on the "playlists" edge with a given conditions (other predicates).
func HasPlaylistsWith(preds ...predicate.Playlist) predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := newPlaylistsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Track) predicate.Track {
	return predicate.Track(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Track) predicate.Track {
	return predicate.Track(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Track) predicate.Track {
	return predicate.Track(sql.NotPredicates(p))
}

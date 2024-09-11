package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CSVRecord struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SeriesReference string             `json:"series_reference" bson:"series_reference" validate:"required"`
	Period          string             `json:"period" bson:"period" validate:"required"`
	DataValue       float64            `json:"data_value" bson:"data_value" validate:"required"`
	Suppressed      string             `json:"suppressed" bson:"suppressed" validate:"required"`
	Status          string             `json:"status" bson:"status" validate:"required"`
	Units           string             `json:"units" bson:"units" validate:"required"`
	Magnitude       int                `json:"magnitude" bson:"magnitude" validate:"required"`
	Subject         string             `json:"subject" bson:"subject" validate:"required"`
	Group           string             `json:"group" bson:"group" validate:"required"`
	SeriesTitle1    string             `json:"series_title_1" bson:"series_title_1" validate:"required"`
	SeriesTitle2    string             `json:"series_title_2" bson:"series_title_2" validate:"required"`
	SeriesTitle3    string             `json:"series_title_3" bson:"series_title_3" validate:"required"`
	SeriesTitle4    string             `json:"series_title_4" bson:"series_title_4" validate:"required"`
	SeriesTitle5    string             `json:"series_title_5" bson:"series_title_5" validate:"required"`
}

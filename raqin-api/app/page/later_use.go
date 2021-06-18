package page

import "raqin-api/utils/validator"

func (bgSrvc *pageService) RevisionsByReviewerID(pageRev ByID) (*[]PageRevisionResponse, error) {

	if err := validator.Validate(pageRev); err != nil {
		return nil, err
	}

	revs, err := bgSrvc.pageRepo.RevisionsByReviewerID(pageRev.ID)
	if err != nil {
		return nil, err
	}

	revResponse := []PageRevisionResponse{}
	for _, revision := range *revs {
		rev := PageRevisionResponse{
			ID:          revision.ID,
			Text:        revision.PageText.String,
			LastUpdated: revision.UpdatedAt,
			PageNum:     revision.R.Page.Number,
			// BookName:    revision.R.Page.R.Book.Name,
			Reviewer: Reviewer{
				Name:   revision.R.Reviewer.FirstName + " " + revision.R.Reviewer.LastName,
				Gender: revision.R.Reviewer.Gender.String,
			},
		}
		revResponse = append(revResponse, rev)
	}

	return &revResponse, nil
}

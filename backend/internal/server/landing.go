package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s Server) handleLandingImages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		images, err := s.contentService.LandingImages(ctx)
		if err != nil {
			writeError(w, err)
			return
		}

		if err := writeJSON(w, http.StatusOK, images); err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleUploadLandingImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		file, _, err := r.FormFile("file")
		if err != nil {
			writeError(w, err)
			return
		}

		id, err := s.contentService.UploadLandingImage(ctx, file)
		if err != nil {
			writeError(w, err)
			return
		}

		img, err := s.contentService.LandingImageByID(ctx, id)
		if err != nil {
			writeError(w, err)
			return
		}

		if err := writeJSON(w, http.StatusOK, img); err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleDeleteLandingImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := strconv.Atoi(chi.URLParam(r, "imageID"))
		if err != nil {
			writeError(w, err)
			return
		}

		err = s.contentService.DeleteLandingImage(ctx, int64(id))
		if err != nil {
			writeError(w, err)
			return
		}
	}
}

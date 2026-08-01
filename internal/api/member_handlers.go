package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"concurrent-task-queue-server/internal/member"
)

type MemberHandler struct {
	Store *member.Store
}

func (h *MemberHandler) HandleMembers(
	w http.ResponseWriter,
	r *http.Request,
) {
	switch r.Method {
	case http.MethodPost:
		h.CreateMember(w, r)

	case http.MethodGet:
		h.ListMembers(w)

	default:
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func (h *MemberHandler) CreateMember(
	w http.ResponseWriter,
	r *http.Request,
) {
	var newMember member.Member

	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		http.Error(
			w,
			"Invalid JSON request body",
			http.StatusBadRequest,
		)
		return
	}

	newMember.Name = strings.TrimSpace(newMember.Name)
	newMember.Contact = strings.TrimSpace(newMember.Contact)
	newMember.Role = strings.TrimSpace(newMember.Role)

	if newMember.Name == "" {
		http.Error(
			w,
			"Member name is required",
			http.StatusBadRequest,
		)
		return
	}

	createdMember := h.Store.Create(newMember)

	writeJSON(
		w,
		http.StatusCreated,
		createdMember,
	)
}

func (h *MemberHandler) ListMembers(
	w http.ResponseWriter,
) {
	members := h.Store.GetAll()

	writeJSON(
		w,
		http.StatusOK,
		members,
	)
}

func (h *MemberHandler) GetMember(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	idString := strings.TrimPrefix(
		r.URL.Path,
		"/members/",
	)

	if idString == "" {
		http.Error(
			w,
			"Member ID is required",
			http.StatusBadRequest,
		)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(
			w,
			"Invalid member ID",
			http.StatusBadRequest,
		)
		return
	}

	foundMember, exists := h.Store.Get(id)
	if !exists {
		http.Error(
			w,
			"Member not found",
			http.StatusNotFound,
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		foundMember,
	)
}

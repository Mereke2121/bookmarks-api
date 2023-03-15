package repository

import "github.com/bookmarks-api/models"

func (r *Repository) GetAllItems(userId int) ([]models.Item, error) {
	var items []models.Item
	query := `select id, url, title from bookmarks_items where user_id=$1`

	err := r.db.Select(&items, query, userId)
	if err != nil {
		return items, err
	}
	return items, nil
}

func (r *Repository) AddItem(item *models.Item) error {
	query := `insert into bookmarks_items (user_id, url, title) values($1, $2, $3)`
	_, err := r.db.Exec(query, item.UserId, item.Url, item.Title)
	return err
}

func (r *Repository) DeleteItem(id, userId int) error {
	query := `delete from bookmarks_items where id=$1 and user_id=$2`
	_, err := r.db.Exec(query, id, userId)
	return err
}

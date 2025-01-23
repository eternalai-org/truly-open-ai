package daos

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FirstTwitterTweetByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.TwitterTweet, error) {
	var m models.TwitterTweet
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstTwitterTweet(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.TwitterTweet, error) {
	var m models.TwitterTweet
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindTwitterTweet(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.TwitterTweet, error) {
	var ms []*models.TwitterTweet
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindTwitterTweet4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.TwitterTweet, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.TwitterTweet
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.TwitterTweet{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}
func (d *DAO) GetListTwitterScan(tx *gorm.DB) ([]string, error) {
	var rs []struct {
		Username string
	}

	query := tx.Raw(`
		select distinct username
		from (
				select twitter_username username
				from agent_infos
				where agent_type = 1
				and scan_enabled = 1
				and agent_contract_id != ''
				and twitter_info_id > 0
				and eai_balance >= 0.1
				union
				select twitter_username username
				from twitter_followings
			) rs
		order by rand()
	`)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}

	// Extract usernames into a slice of strings
	usernames := make([]string, len(rs))
	for i, record := range rs {
		usernames[i] = record.Username
	}

	return usernames, nil
}

func (d *DAO) GetListTwitterMentionsScan(tx *gorm.DB) ([]string, error) {
	var rs []struct {
		Username string
	}

	query := tx.Raw(`
		select username
		from twitter_scans
		where enabled = 1 and is_mention=1 and scanned = 0
	`)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}

	// Extract usernames into a slice of strings
	usernames := make([]string, len(rs))
	for i, record := range rs {
		usernames[i] = record.Username
	}

	return usernames, nil
}

func (d *DAO) GetListTwitterDefaultFollow(tx *gorm.DB) ([]string, error) {
	var rs []struct {
		ID string
	}

	query := tx.Raw(`
		select twitter_id id
		from twitter_scans
		where enabled = 1 and is_mention=0
		order by id asc
	`)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}

	ids := make([]string, len(rs))
	for i, record := range rs {
		ids[i] = record.ID
	}

	return ids, nil
}

func (d *DAO) IsTweetReplied(tx *gorm.DB, tweetID string) (bool, error) {
	var rs []struct {
		ID string
	}

	query := tx.Raw(`
		select tweetid id 
		from agent_snapshot_post_actions aspa 
		where 1=1
		and status ='done'
		and type = 'reply'
		and tweetid = ?
	`, tweetID)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, errs.NewError(err)
	}

	if len(rs) > 0 {
		return true, nil
	}

	return false, nil
}

package model

import (
	"context"
	"database/sql"
	"mt_node/internal/types"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChainInfo struct {
	Id                    int64          `db:"id"`
	ChainId               string         `db:"chain_id"`
	ChainIdDecimal        int64          `db:"chain_id_decimal"`
	CreatedAt             time.Time      `db:"created_at"`
	UpdatedAt             time.Time      `db:"updated_at"`
	Name                  string         `db:"name"`
	ShortName             sql.NullString `db:"short_name"`
	Slug                  sql.NullString `db:"slug"`
	ChainType             sql.NullString `db:"chain_type"`
	Environment           sql.NullString `db:"environment"`
	IsActive              bool           `db:"is_active"`
	RpcUrls               sql.NullString `db:"rpc_urls"`
	WssUrls               sql.NullString `db:"wss_urls"`
	DefaultRpcIndex       int64          `db:"default_rpc_index"`
	ExplorerUrl           sql.NullString `db:"explorer_url"`
	ApiUrl                sql.NullString `db:"api_url"`
	NativeCurrency        sql.NullString `db:"native_currency"`
	IconUrl               sql.NullString `db:"icon_url"`
	Color                 sql.NullString `db:"color"`
	IsTestnet             bool           `db:"is_testnet"`
	CurrencySymbol        sql.NullString `db:"currency_symbol"`
	BlockExplorerName     sql.NullString `db:"block_explorer_name"`
	DisplayOrder          int64          `db:"display_order"`
	IsDeprecated          bool           `db:"is_deprecated"`
	SupportsEip1559       bool           `db:"supports_eip1559"`
	SupportsBatchRequests bool           `db:"supports_batch_requests"`
}

func (m *ChainInfo) TableName() string {
	return "mt_auth.chain_info"
}

type ChainInfoModel struct {
	db *pgxpool.Pool
}

func NewChainInfoModel(db *pgxpool.Pool) *ChainInfoModel {
	return &ChainInfoModel{db: db}
}

func (m *ChainInfoModel) FindAll(ctx context.Context) ([]*types.ChainInfo, error) {
	rows, err := m.db.Query(ctx, "SELECT * FROM mt_auth.chain_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dbRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[ChainInfo])
	if err != nil {
		return nil, err
	}

	chainInfos := make([]*types.ChainInfo, 0, len(dbRows))
	for i := range dbRows {
		chainInfos = append(chainInfos, dbRows[i].ToTypes())
	}

	return chainInfos, nil
}

func (m *ChainInfo) ToTypes() *types.ChainInfo {
	return &types.ChainInfo{
		Id:                    m.Id,
		ChainId:               m.ChainId,
		ChainIdDecimal:        m.ChainIdDecimal,
		Name:                  m.Name,
		ShortName:             m.ShortName.String,
		Slug:                  m.Slug.String,
		ChainType:             m.ChainType.String,
		Environment:           m.Environment.String,
		IsActive:              m.IsActive,
		RpcUrls:               m.RpcUrls.String,
		WssUrls:               m.WssUrls.String,
		DefaultRpcIndex:       m.DefaultRpcIndex,
		ExplorerUrl:           m.ExplorerUrl.String,
		ApiUrl:                m.ApiUrl.String,
		NativeCurrency:        m.NativeCurrency.String,
		IconUrl:               m.IconUrl.String,
		Color:                 m.Color.String,
		IsTestnet:             m.IsTestnet,
		CurrencySymbol:        m.CurrencySymbol.String,
		BlockExplorerName:     m.BlockExplorerName.String,
		DisplayOrder:          m.DisplayOrder,
		IsDeprecated:          m.IsDeprecated,
		SupportsEip1559:       m.SupportsEip1559,
		SupportsBatchRequests: m.SupportsBatchRequests,
		CreatedAt:             m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:             m.UpdatedAt.Format(time.RFC3339),
	}
}

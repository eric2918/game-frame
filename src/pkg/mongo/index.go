package mongo

func InitIndex() {
	// login db
	UniqueIndex(ACCOUNT_DB, ACCOUNTS_COLLECTION, []string{"username"})

	// game db
	UniqueIndex(GAME_DB, PLAYERS_COLLECTION, []string{"playerId"})
	UniqueIndex(GAME_DB, PLAYERS_COLLECTION, []string{"accountId"})

	//UniqueIndex(GAME_DB, ROLES_COLLECTION, []string{"roleId"})
	//
	//UniqueIndex(GAME_DB, SKILLS_COLLECTION, []string{"skillId"})
	//
	//UniqueIndex(GAME_DB, SKINS_COLLECTION, []string{"roleId", "skinId"})
	//
	//UniqueIndex(GAME_DB, PROPS_COLLECTION, []string{"propId"})
}

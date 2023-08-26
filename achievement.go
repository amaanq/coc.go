package coc

import "errors"

type AchievementID string

type Achievement struct {
	Name           string `json:"name"`
	Info           string `json:"info"`
	CompletionInfo string `json:"completionInfo"`
	Village        string `json:"village"`
	Stars          int    `json:"stars"`
	Value          int    `json:"value"`
	Target         int    `json:"target"`
}

// If you're peering into this and I've typo'd an achievement name causing an error when fetching it, please let me know.
var (
	None                    Achievement = Achievement{}
	BiggerCoffers           Achievement = Achievement{Name: "Bigger Coffers"}
	GetThoseGoblins         Achievement = Achievement{Name: "Get those Goblins!"}
	BiggerAndBetter         Achievement = Achievement{Name: "Bigger & Better"}
	NiceAndTidy             Achievement = Achievement{Name: "Nice and Tidy"}
	DiscoverNewTroops       Achievement = Achievement{Name: "Discover New Troops"}
	GoldGrab                Achievement = Achievement{Name: "Gold Grab"}
	ElixirEscapade          Achievement = Achievement{Name: "Elixir Escapade"}
	SweetVictory            Achievement = Achievement{Name: "Sweet Victory!"}
	EmpireBuilder           Achievement = Achievement{Name: "Empire Builder"}
	WallBuster              Achievement = Achievement{Name: "Wall Buster"}
	Humiliator              Achievement = Achievement{Name: "Humiliator"}
	UnionBuster             Achievement = Achievement{Name: "Union Buster"}
	Conqueror               Achievement = Achievement{Name: "Conqueror"}
	Unbreakable             Achievement = Achievement{Name: "Unbreakable"}
	FriendInNeed            Achievement = Achievement{Name: "Friend in Need"}
	MortarMauler            Achievement = Achievement{Name: "Mortar Mauler"}
	HeroicHeist             Achievement = Achievement{Name: "Heroic Heist"}
	LeagueAllStar           Achievement = Achievement{Name: "League All-Star"}
	XBowExterminator        Achievement = Achievement{Name: "X-Bow Exterminator"}
	Firefighter             Achievement = Achievement{Name: "Firefighter"}
	WarHero                 Achievement = Achievement{Name: "War Hero"}
	ClanWarWealth           Achievement = Achievement{Name: "Clan War Wealth"}
	AntiArtillery           Achievement = Achievement{Name: "Anti-Artillery"}
	SharingIsCaring         Achievement = Achievement{Name: "Sharing is caring"}
	KeepYourAccountSafeOld  Achievement = Achievement{Name: "Keep Your Account Safe!", Info: "Protect your village by connecting to a social network"}
	MasterEngineering       Achievement = Achievement{Name: "Master Engineering"}
	NextGenerationModel     Achievement = Achievement{Name: "Next Generation Model"}
	UnBuildIt               Achievement = Achievement{Name: "Un-Build It"}
	ChampionBuilder         Achievement = Achievement{Name: "Champion Builder"}
	HighGear                Achievement = Achievement{Name: "High Gear"}
	HiddenTreasures         Achievement = Achievement{Name: "Hidden Treasures"}
	GamesChampion           Achievement = Achievement{Name: "Games Champion"}
	DragonSlayer            Achievement = Achievement{Name: "Dragon Slayer"}
	WarLeagueLegend         Achievement = Achievement{Name: "War League Legend"}
	KeepYourAccountSafeSCID Achievement = Achievement{Name: "Keep Your Account Safe!", Info: "Connect your account to Supercell ID for safe keeping."}
	WellSeasoned            Achievement = Achievement{Name: "Well Seasoned"}
	ShatteredAndScattered   Achievement = Achievement{Name: "Shattered and Scattered"}
	NotSoEasyThisTime       Achievement = Achievement{Name: "Not So Easy This Time"}
	BustThis                Achievement = Achievement{Name: "Bust This"}
	SuperbWork              Achievement = Achievement{Name: "Superb Work"}
	SiegeSharer             Achievement = Achievement{Name: "Siege Sharer"}
	Counterspell            Achievement = Achievement{Name: "Counterspell"}
	MonolithMasher          Achievement = Achievement{Name: "Monolith Masher"}
	GetThoseOtherGoblins    Achievement = Achievement{Name: "Get those other Goblins!"}
	GetEvenMoreGoblins      Achievement = Achievement{Name: "Get even more Goblins!"}
	UngratefulChild         Achievement = Achievement{Name: "Ungrateful Child"}
	AggressiveCapitalism    Achievement = Achievement{Name: "Aggressive Capitalism"}
	MostValuableClanmate    Achievement = Achievement{Name: "Most Valuable Clanmate"}
)

func GetAchievement(givenAchievements []Achievement, wantedAchievement Achievement) (Achievement, error) {
	for _, achievement := range givenAchievements {
		switch wantedAchievement {
		case KeepYourAccountSafeOld, KeepYourAccountSafeSCID:
			if achievement.Name == wantedAchievement.Name && achievement.Info == wantedAchievement.Info {
				return achievement, nil
			}
		default:
			if achievement.Name == wantedAchievement.Name {
				return achievement, nil
			}
		}
	}
	return None, errors.New("no matching achievement found")
}

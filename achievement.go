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
	KeepYourAccountSafeOld  Achievement = Achievement{Name: "Keep Your Account Safe!", Info: "Protect your village by connecting to a social network"}
	KeepYourAccountSafeSCID Achievement = Achievement{Name: "Keep Your Account Safe!", Info: "Connect your account to Supercell ID for safe keeping."}
	BiggerAndBetter         Achievement = Achievement{Name: "Bigger & Better"}
	DiscoverNewTroops       Achievement = Achievement{Name: "Discover New Troops"}
	BiggerCoffers           Achievement = Achievement{Name: "Bigger Coffers"}
	GoldGrab                Achievement = Achievement{Name: "Gold Grab"}
	ElixirEscapade          Achievement = Achievement{Name: "Elixir Escapade"}
	HeroicHeist             Achievement = Achievement{Name: "Heroic Heist"}
	WellSeasoned            Achievement = Achievement{Name: "Well Seasoned"}
	NiceAndTidy             Achievement = Achievement{Name: "Nice and Tidy"}
	EmpireBuilder           Achievement = Achievement{Name: "Empire Builder"}
	ClanWarWealth           Achievement = Achievement{Name: "Clan War Wealth"}
	FriendInNeed            Achievement = Achievement{Name: "Friend in Need"}
	SharingIsCaring         Achievement = Achievement{Name: "Sharing is caring"}
	SiegeSharer             Achievement = Achievement{Name: "Siege Sharer"}
	WarHero                 Achievement = Achievement{Name: "War Hero"}
	WarLeagueLegend         Achievement = Achievement{Name: "War League Legend"}
	GamesChampion           Achievement = Achievement{Name: "Games Champion"}
	Unbreakable             Achievement = Achievement{Name: "Unbreakable"}
	SweetVictory            Achievement = Achievement{Name: "Sweet Victory!"}
	Conqueror               Achievement = Achievement{Name: "Conqueror"}
	LeagueAllStar           Achievement = Achievement{Name: "League All-Star"}
	Humiliator              Achievement = Achievement{Name: "Humiliator"}
	NotSoEasyThisTime       Achievement = Achievement{Name: "Not So Easy This Time"}
	UnionBuster             Achievement = Achievement{Name: "Union Buster"}
	BustThis                Achievement = Achievement{Name: "Bust This"}
	WallBuster              Achievement = Achievement{Name: "Wall Buster"}
	MortarMauler            Achievement = Achievement{Name: "Mortar Mauler"}
	XBowExterminator        Achievement = Achievement{Name: "X-Bow Exterminator"}
	Firefighter             Achievement = Achievement{Name: "Firefighter"}
	AntiArtillery           Achievement = Achievement{Name: "Anti-Artillery"}
	ShatteredAndScattered   Achievement = Achievement{Name: "Shattered and Scattered"}
	Counterspell            Achievement = Achievement{Name: "Counterspell"}
	MonolithMasher          Achievement = Achievement{Name: "Monolith Masher"}
	GetThoseGoblins         Achievement = Achievement{Name: "Get those Goblins!"}
	GetThoseOtherGoblins    Achievement = Achievement{Name: "Get those other Goblins!"}
	GetEvenMoreGoblins      Achievement = Achievement{Name: "Get even more Goblins!"}
	DragonSlayer            Achievement = Achievement{Name: "Dragon Slayer"}
	UngratefulChild         Achievement = Achievement{Name: "Ungrateful Child"}
	SuperbWork              Achievement = Achievement{Name: "Superb Work"}

	MasterEngineering   Achievement = Achievement{Name: "Master Engineering"}
	HiddenTreasures     Achievement = Achievement{Name: "Hidden Treasures"}
	HighGear            Achievement = Achievement{Name: "High Gear"}
	NextGenerationModel Achievement = Achievement{Name: "Next Generation Model"}
	UnBuildIt           Achievement = Achievement{Name: "Un-Build It"}
	ChampionBuilder     Achievement = Achievement{Name: "Champion Builder"}

	AggressiveCapitalism Achievement = Achievement{Name: "Aggressive Capitalism"}
	MostValuableClanmate Achievement = Achievement{Name: "Most Valuable Clanmate"}
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

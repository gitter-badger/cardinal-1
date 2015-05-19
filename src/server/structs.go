package main

import "encoding/json"

type User struct {
    Username string
    Password string
    DashItems []DashItem
    Collections []Collection
}

type Legality struct {
    Modern string
    Legacy string
    Vintage string
    Freeform string
    Prismatic string
    TribalWarsLegacy string
    Singleton string
    Commander string
}

type ForeignName struct {
    Language string
    Name string
}

type HearthStoneCard struct {
    Name string
    Cost int
    Type string
    Rarity string
    Faction string
    Text string
    Mechanics []string
    Flavor string
    Artist string
    Attack int
    Health int
    Collectible bool
    Id string
    Elite bool
    PlayerClass string
    InPlayText string
    Durability int
    HowToGet string
    HowToGetGold string
    Faction string
}

type MagicCard struct {
    Id string
    Layout string
    Name string
    ManaCost string
    Cmc int
    Colors []string
    Type string
    Types []string
    SubTypes []string
    Power int
    Toughness int
    Text string
    ForeignNames []ForeignName
    Printings []string
    Legalities Legality
    ImageNames []string
}

type Collection struct {
    Name string
    Game string
    IsMain bool
    Owner User
    Cards []interface{}
}

type DashItem struct {
    Img string
    Title string
    Content string
}

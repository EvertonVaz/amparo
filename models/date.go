package models

type ImportantDates struct {
    ID                  int    `json:"id"`
    SeventhDayMass      string `json:"seventhDayMass"`
    DeathRegistration   string `json:"deathRegistration"`
    DeathPensionRequest string `json:"deathPensionRequest"`
    InventoryOpening    string `json:"inventoryOpening"`
    LifeInsuranceClaim  string `json:"lifeInsuranceClaim"`
}

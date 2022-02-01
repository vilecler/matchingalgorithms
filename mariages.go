package main

import (
	"fmt"
)

// Objet Mariages
//
// Map qui permet de stocker l'ensemble des appariements

type Mariages map[AgentID]AgentID

func (m Mariages) Debug(){
  for proposantID, disposantID := range m{
    fmt.Println(proposantID, "is married with", disposantID, ".")
  }
}

/*
 * Fonction IsStable
 *
 * Permet d'évaluer si un ensemble d'appariement est stable ou non.
 */
func (m Mariages) IsStable(proposants []Agent, disposants []Agent) bool{
  for proposantID := range m{
    for proposantID2  := range m{
      disposantID := m[proposantID]
      disposantID2 := m[proposantID2]

      if(proposantID != proposantID2 && disposantID != disposantID2){
        proposant := GetAgent(proposants, proposantID)
        disposant := GetAgent(disposants, disposantID)
        proposant2 := GetAgent(proposants, proposantID2)
        disposant2 := GetAgent(disposants, disposantID2)

        if proposant.Prefers(disposant2, disposant) && disposant2.Prefers(proposant, proposant2){ //le mariage n'est pas stable
          return false
        }
        if disposant.Prefers(proposant2, proposant) && proposant2.Prefers(disposant, disposant2){ //le mariage n'est pas stable
          return false
        }
      }
    }
  }
  return true
}

func (m mariages) quantifyBias(agtA []*Agent, agtB []*Agent) (Poidsprop int, Poidsdisp int){
  S_prop := 0
  S_disp := 0
  for key, val := range(a) {
    prop := GetAgentById(key, agtA)
    disp := GetAgentById(val, agtB)
    sp, sd := 0,0
    preference_prop := prop.Prefs[0]
    preference_disp := disp.Prefs[0]
    c_disp, c_prop := 0,0
    for val != preference_prop {
      sp++
      c_prop++
      preference_prop = prop.Prefs[c_prop]
      if val == preference_prop{
        break
      }
    }
    for key !=  preference_disp{
      sd++
      c_disp++
      preference_disp = disp.Prefs[c_disp]
      if key == preference_disp{
        break
      }
    }
    S_prop = S_prop + sp
    S_disp = S_disp + sd
  }
  return S_prop, S_disp

}

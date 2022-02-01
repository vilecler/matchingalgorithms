package main;

import(
  "fmt"
)

// Objet Propositions
//
// Map qui permet de stocker des propositions que font les proposants aux disposants.

type Propositions map[AgentID][]AgentID

func (p Propositions) Debug(){
  for disposantID, proposants := range p{
    fmt.Println(disposantID, "has received propositions from", proposants, ".")
  }
}

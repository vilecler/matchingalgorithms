package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type AgentID string

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []AgentID
	iterateur int
}

func NewAgent(id AgentID, name string, prefs []AgentID) Agent {
	return Agent{
		id,
		name,
		prefs,
		0,
	}
}

func Equal(ag1 Agent, ag2 Agent) bool {
	if ag1.ID != ag2.ID {
		return false
	}

	// Pas obligatoire à partir de là, à discuter...
	if ag1.Name != ag2.Name {
		return false
	}

	if len(ag1.Prefs) != len(ag2.Prefs) {
		return false
	}

	for i := range ag1.Prefs {
		if ag1.Prefs[i] != ag1.Prefs[i] {
			return false
		}
	}

	return true
}

func (a Agent) String() string {
	return fmt.Sprintf("%s %s %v", a.ID, a.Name, a.Prefs)
}

func (a Agent) rank(b Agent) (int, error) {
	for i, v := range a.Prefs {
		if v == b.ID {
			return i, nil
		}
	}
	return -1, errors.New("Agent %s not found" + string(b.ID))
}

// renvoie vrai si ag préfère a à b
func (ag Agent) Prefers(a, b Agent) bool {
	r1, err1 := ag.rank(a)
	if err1 != nil {
		return false
	}

	r2, err2 := ag.rank(b)
	if err2 != nil {
		return false
	}

	return r1 < r2
}

func RandomPrefs(ids []AgentID) (res []AgentID) {
	res = make([]AgentID, len(ids))
	copy(res, ids)
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	return
}

func (agent *Agent) iteration() {
  agent.iterateur ++
}

func (agent *Agent) PrefIterator () (id  AgentID){
  pref := agent.Prefs[0]
  if len(agent.Prefs) > 1{
    agent.Prefs = agent.Prefs[1:]
  }
  return pref
}

/*
 *  Fonction GenerateAgents
 *
 *	Permet de générer les deux groupes d'agents à partir d'une liste de prénom et d'une taille n.
 *
 */
func GenerateAgents(prenoms []Prenom, n int) (proposants []Agent, disposants []Agent){
	males := ShufflePrenoms(GetAllMales(prenoms), n)	//prénoms masculins
	females := ShufflePrenoms(GetAllFemales(prenoms), n)	//prénoms féminins

	proposants = make([]Agent, 0, n)
	disposants = make([]Agent, 0, n)

	proposantPrefix := "p"
	disposantPrefix := "d"

	prefsProposants := make([]AgentID, n)
	prefsDisposants := make([]AgentID, n)

	for i := 0; i < n; i++ {	//generating preferences
		prefsProposants[i] = AgentID(proposantPrefix + fmt.Sprintf("%d", i))
		prefsDisposants[i] = AgentID(disposantPrefix + fmt.Sprintf("%d", i))
	}

	for i := 0; i < n; i++ {	//generating agent
		prefsProposant := RandomPrefs(prefsDisposants)
		proposant := NewAgent(prefsProposants[i], males[i].prenom, prefsProposant)
		proposants = append(proposants, proposant)

		prefsDisposant := RandomPrefs(prefsProposants)
		disposant := NewAgent(prefsDisposants[i], females[i].prenom, prefsDisposant)
		disposants = append(disposants, disposant)
	}

	return proposants, disposants
}

func GenerateProblematicAgentsForDynamiqueLibreAlgorithm(prenoms []Prenom) (proposants []Agent, disposants []Agent){
	males := ShufflePrenoms(GetAllMales(prenoms), 3)	//prénoms masculins
	females := ShufflePrenoms(GetAllFemales(prenoms), 3)	//prénoms féminins

	proposants = make([]Agent, 0, 3)
	disposants = make([]Agent, 0, 3)

	proposants = append(proposants, NewAgent("p0", males[0].prenom, []AgentID{"d0", "d1", "d2"}))
	proposants = append(proposants, NewAgent("p1", males[1].prenom, []AgentID{"d0", "d2", "d1"}))
	proposants = append(proposants, NewAgent("p2", males[2].prenom, []AgentID{"d2", "d0", "d1"}))

	disposants = append(disposants, NewAgent("d0", females[0].prenom, []AgentID{"p1", "p0", "p2"}))
	disposants = append(disposants, NewAgent("d1", females[1].prenom, []AgentID{"p2", "p0", "p1"}))
	disposants = append(disposants, NewAgent("d2", females[2].prenom, []AgentID{"p1", "p2", "p0"}))


	return proposants, disposants
}

/*
 *  Fonction GenerateProblematicAgentsForAcceptationImmediateAlgorithm
 *
 *	Renvoie une instance du problème d'appariement pour laquelle l'algorithme d'acceptation immédiate proposera un appariement instable.
 *
 */
func GenerateProblematicAgentsForAcceptationImmediateAlgorithm(prenoms []Prenom) (proposants []Agent, disposants []Agent){
	males := ShufflePrenoms(GetAllMales(prenoms), 3)	//prénoms masculins
	females := ShufflePrenoms(GetAllFemales(prenoms), 3)	//prénoms féminins

	proposants = make([]Agent, 0, 3)
	disposants = make([]Agent, 0, 3)

	proposants = append(proposants, NewAgent("p0", males[0].prenom, []AgentID{"d0", "d1", "d2"}))
	proposants = append(proposants, NewAgent("p1", males[1].prenom, []AgentID{"d0", "d1", "d2"}))
	proposants = append(proposants, NewAgent("p2", males[2].prenom, []AgentID{"d1", "d0", "d2"}))

	disposants = append(disposants, NewAgent("d0", females[0].prenom, []AgentID{"p0", "p1", "p2"}))
	disposants = append(disposants, NewAgent("d1", females[1].prenom, []AgentID{"p1", "p0", "p2"}))
	disposants = append(disposants, NewAgent("d2", females[2].prenom, []AgentID{"p0", "p2", "p1"}))


	return proposants, disposants
}

/*
 * Fonction TableauAgTOtableauPointeursAg
 *
 * Transforme un tableau d'Agents en tableau de pointeurs d'Agents
 */
func TableauAgTOtableauPointeursAg(tab []Agent) []*Agent {
  taille := len(tab)
  tabbis := make([]*Agent, taille)
  for i, _ := range(tab){
    tabbis[i] = &tab[i]
  }
  return tabbis
}

/*
 *  Fonction GetAgent
 *
 *	Permet de récupérer un agent à partir de son identifiant dans une liste d'agents.
 *
 */
func GetAgent(agents []Agent, ID AgentID) Agent{
	for i := 0; i < len(agents); i++{
		if(agents[i].ID == ID){
			return agents[i]
		}
	}
	return Agent{}
}

func GetAgentById(id AgentID, tab []*Agent)  (x *Agent){
  for _,agent := range(tab){
    if agent.ID == id{
      return agent
    }
  }
  return nil
}


/*
 * Fonction RemoveIndex
 *
 * Supprime le ième Agent dans un tableau d'agents
 */
func RemoveIndex(s []Agent, index int) []Agent {
 var r []Agent
 for i := 0; i < len(s); i++{
	 if i != index{
		 r = append(r, s[i])
	 }
 }
 return r
}

/*
* Fonction RemovePref
*
* Supprime la ième Préférence des préférences d'un agent
*/
func RemovePrefs(prefs []AgentID, index int) []AgentID{
 var r []AgentID
 for i := 0; i < len(prefs); i++{
	 if i != index{
		 r = append(r, prefs[i])
	 }
 }
 return r
}

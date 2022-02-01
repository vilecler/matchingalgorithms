package main

// Objet Configuration
//
// Permet de stocker une instance du problème pour n fixé

type Configuration struct{
  proposants []Agent
  disposants []Agent
}

func NewConfiguration(proposants []Agent, disposants []Agent) Configuration{
  return Configuration{
    proposants,
    disposants,
  }
}

func (c Configuration) Equal(conf Configuration) bool{

  if len(c.proposants) != len(conf.proposants){
    return false
  }

  //compare proposants
  for i := 0; i < len(c.proposants); i++{
    if c.proposants[i].ID != conf.proposants[i].ID{
      return false
    }

    for j := 0; j < len(c.proposants[i].Prefs); j++{
      if c.proposants[i].Prefs[j] != conf.proposants[i].Prefs[j]{
        return false
      }
    }
  }

  //compare disposants
  for i := 0; i < len(c.disposants); i++{
    if c.disposants[i].ID != conf.disposants[i].ID{
      return false
    }

    for j := 0; j < len(c.disposants[i].Prefs); j++{
      if c.disposants[i].Prefs[j] != conf.disposants[i].Prefs[j]{
        return false
      }
    }
  }

  return true
}

func AddConfiguration(configurations []Configuration, c Configuration) []Configuration{
  if !CanAddConfiguration(configurations, c){
    return configurations
  }

  configurations = append(configurations, c)
  return configurations
}

func CanAddConfiguration(configurations []Configuration, c Configuration) bool{
  for i := 0; i < len(configurations); i++{
    if configurations[i].Equal(c) {
      return false
    }
  }
  return true
}

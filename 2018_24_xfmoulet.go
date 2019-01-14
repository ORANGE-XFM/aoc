// Go version but with direct data, no more reading data, just getting them into shape with the keyboard

package main

import (
	"fmt"
	"sort"
)

const (
    normal = iota
    weak
    immune
)

const (
	radiation = iota
	bludgeoning
	fire 
	cold
	slashing
)

const (
	s_immune = iota
	s_infection
)

type Group struct {
	system, id int 
	units, hitpoints int
	attack_pts, attack int
	initiative int
	mods [5]byte
	//
	defending *Group
	attacked bool
}

var test_systems = [] Group { 
		{ system : s_immune,    id: 1, units:17,   hitpoints:5390, attack_pts:4507, attack:fire,         initiative:2, mods: [5]byte { radiation:weak, bludgeoning:weak } },
		{ system : s_immune,    id: 2, units:989,  hitpoints:1274, attack_pts:25,   attack:slashing,     initiative:3, mods: [5]byte { fire:immune, bludgeoning:weak, slashing:weak } },
		{ system : s_infection, id: 1, units:801,  hitpoints:4706, attack_pts:116,  attack:bludgeoning,  initiative:1, mods: [5]byte { radiation:weak } },
		{ system : s_infection, id: 2, units:4485, hitpoints:2961, attack_pts:12,   attack:slashing,     initiative:4, mods: [5]byte { radiation:immune,fire:weak, cold:weak } },
	}

var systems = [] Group {
		{ system : s_immune,    units:2208, hitpoints: 6238,  attack_pts:23,  attack:bludgeoning, initiative:20, mods: [5]byte { slashing:immune } },
		{ system : s_immune,    units:7603, hitpoints: 6395,  attack_pts:6,   attack:cold,        initiative:15, mods: [5]byte { radiation:weak } },
		{ system : s_immune,    units:4859, hitpoints: 5904,  attack_pts:12,  attack:cold,        initiative:11, mods: [5]byte { fire:weak } },
		{ system : s_immune,    units:1608, hitpoints: 7045,  attack_pts:31,  attack:radiation,   initiative:10, mods: [5]byte { fire:weak, cold:weak,  bludgeoning:immune, radiation:immune } },
		{ system : s_immune,    units:39,   hitpoints: 4208,  attack_pts:903, attack:radiation,   initiative:7 },
		{ system : s_immune,    units:6969, hitpoints: 9562,  attack_pts:13,  attack:slashing,    initiative:3,  mods: [5]byte { slashing:immune, cold:immune } },
		{ system : s_immune,    units:2483, hitpoints: 6054,  attack_pts:20,  attack:cold,        initiative:19, mods: [5]byte { fire:immune } },
		{ system : s_immune,    units:506,  hitpoints: 3336,  attack_pts:64,  attack:radiation,   initiative:6 },
		{ system : s_immune,    units:2260, hitpoints: 10174, attack_pts:34,  attack:slashing,    initiative:5,  mods: [5]byte { fire:weak } },
		{ system : s_immune,    units:2817, hitpoints: 9549,  attack_pts:31,  attack:cold,        initiative:2,  mods: [5]byte { cold:immune, fire:immune, bludgeoning:weak } },
		{ system : s_infection, units:3650, hitpoints: 25061, attack_pts:11,  attack:slashing,    initiative:12, mods: [5]byte { fire:weak, bludgeoning:weak } },
		{ system : s_infection, units:508,  hitpoints: 48731, attack_pts:172, attack:cold,        initiative:13, mods: [5]byte { bludgeoning:weak } },
		{ system : s_infection, units:724,  hitpoints: 27385, attack_pts:69,  attack:radiation,   initiative:1 },
		{ system : s_infection, units:188,  hitpoints: 41786, attack_pts:416, attack:bludgeoning, initiative:4 },
		{ system : s_infection, units:3045, hitpoints: 36947, attack_pts:24,  attack:slashing,    initiative:9,  mods: [5]byte { slashing:weak, fire:immune, bludgeoning:immune } },
		{ system : s_infection, units:7006, hitpoints: 42545, attack_pts:9,   attack:fire,        initiative:16, mods: [5]byte { cold:immune, slashing:immune, fire:immune } },
		{ system : s_infection, units:853,  hitpoints: 55723, attack_pts:114, attack:bludgeoning, initiative:17, mods: [5]byte { cold:weak, fire:immune } },
		{ system : s_infection, units:3268, hitpoints: 43027, attack_pts:25,  attack:slashing,    initiative:8,  mods: [5]byte { slashing:immune, fire:immune } },
		{ system : s_infection, units:1630, hitpoints: 47273, attack_pts:57,  attack:slashing,    initiative:14, mods: [5]byte { cold:weak, bludgeoning:weak } },
		{ system : s_infection, units:3383, hitpoints: 12238, attack_pts:7,   attack:radiation,   initiative:18 },
	}

func (att Group) damage(target Group) int {
	var damages = [] int {normal: 1, immune: 0, weak: 2}
	return damages[target.mods[att.attack]]*att.effective_power()
}

func (grp Group) effective_power() int {
	return grp.attack_pts*grp.units
}

func(grp Group) system_name() string {
	return []string{"Immune","Infection"}[grp.system]
}

func turn(groups []Group) int {
	fmt.Println(" --------- start turn --")

	for i, grp := range groups {
		if grp.units == 0 { continue }
		groups[i].attacked = false
		groups[i].defending = nil
		fmt.Println("Group",grp.id,grp.system_name(), grp.units,"units of effective power", grp.effective_power())
	}

	fmt.Println(" -- selection phase --")
	// sort groups by effective power
	sort.Slice(groups, func(i, j int) bool {
			// decreasing eff power, decreasing initiative
			return groups[i].effective_power()*1000+groups[i].initiative > groups[j].effective_power()*1000+groups[j].initiative 
	})

	for att_id,att := range groups {
		if att.units==0 {
			continue // skip killed groups ?
		}
		// get the best attack
		var most_damage int = 0
		var defending *Group  = nil
		for def_id, def := range groups {
			damage := att.damage(def)

			// cant attack those
			if def.system == att.system || damage == 0 || def.attacked || def.units==0 {
				continue
			}

			if damage > most_damage || 
				(damage == most_damage && def.effective_power()>defending.effective_power()) || 
				(damage == most_damage && def.effective_power()==defending.effective_power() && def.initiative > defending.initiative)  { // check ties order : target with the most effective power or initiative
				most_damage = damage
				defending = &groups[def_id]
			}
		}
		if defending != nil {
			fmt.Println(att.system_name(),"group",att.id,"would deal defending group",defending.id,"with damage",most_damage)
			defending.attacked = true
			groups[att_id].defending = defending
		} else {
			fmt.Println(att.system_name(),"group",att.id,"would not attack")
		}
	}

	fmt.Println(" -- attack phase --")

	tmp := make([]*Group, len(groups))
	for i,_ := range groups {
		tmp[i] = &groups[i]
	}

	one_attacked := false
	// sort groups by initiative, decreasing
	sort.Slice(tmp, func(i,j int) bool {
		return tmp[i].initiative > tmp[j].initiative
		})
	for _,att := range tmp {
		if att.units==0 || att.defending == nil {
			continue // skip killed groups ?
		}
		one_attacked = true // check none attacks

		def := att.defending
		units := att.damage(*def) / def.hitpoints
		fmt.Println(att.system_name(),"Group",att.id,"attacks defending group",def.id,"would kill",units,"units over",def.units)
		def.units -= units
		if def.units<0 {
			def.units = 0
		}
	}


	// check winning 
	units := [2]int{0,0}
	for _,grp := range groups {
		units[grp.system] += grp.units
	}

	switch {
		case units[s_immune]==0 && units[s_infection]==0 : return -1 // draw is a defeat 
		case units[s_immune]==0 : return -units[s_infection]
		case units[s_infection]==0 : return units[s_immune]
		case !one_attacked : return -1 // blocked, defeat
		default: return 0
	}
}

func battle(system [] Group, boost int ) int {
	// make a copy ! 
	tmp := make ([]Group, len(system))
	copy(tmp,system)

	// boost it ! 
	for i := range tmp {
		if tmp[i].system ==  s_immune {
			tmp[i].attack_pts += boost
		}
	}

	for {
		res := turn(tmp)
		if res!=0 {
			return res
			break
		}		
	}
	return 0
}

func find_boost(system[] Group) {
	for boost:=0;;boost++ {
		res := battle(system,boost)
		fmt.Println("====== Result of fight with Boost",boost," : ",res)
		if res>0 { return }
	}
}

func main() {
	fmt.Println("====== Result of fight with Boost",0," : ",battle(test_systems,0),"\n")
	fmt.Println("====== Result of fight with Boost",1570," : ",battle(test_systems,1570),"\n")
	find_boost(systems)
}
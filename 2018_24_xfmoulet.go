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
	system, units, hitpoints int
	attack_pts, attack int
	initiative int
	mods [5]byte
	//
	attacks, attacked_by int
}

var test_systems = [] Group { 
		{ system : s_immune,    units:17,   hitpoints:5390, attack_pts:4507, attack:fire,         initiative:2, mods: [5]byte { radiation:weak, bludgeoning:weak } },
		{ system : s_immune,    units:989,  hitpoints:1274, attack_pts:25,   attack:slashing,     initiative:3, mods: [5]byte { fire:immune, bludgeoning:weak, slashing:weak } },
		{ system : s_infection, units:801,  hitpoints:4706, attack_pts:116,  attack:bludgeoning,  initiative:1, mods: [5]byte { radiation:weak } },
		{ system : s_infection, units:4485, hitpoints:2961, attack_pts:12,   attack:slashing,     initiative:4, mods: [5]byte { radiation:immune,fire:weak, cold:weak } },
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

func prepare_attack(groups []Group) {

	// sort groups by effective power
	sort.Slice(groups, func(i, j int) bool {
			return groups[i].effective_power()*1000+groups[i].initiative > groups[j].effective_power()*1000+groups[j].initiative // decreasing eff power, decreasing initiative
	})

	for i:=0;i<len(groups);i++ {
		groups[i].attacked_by = -1
		groups[i].attacks = -1
	}

	for att_id,att := range groups {
		// get the best attack
		var most_damage int = 0
		attacked := 0
		for def_id, def := range groups {
			damage := att.damage(def)
			if def.system == att.system || damage == 0 || def.attacked_by!=-1 {
				continue
			}

			if damage > most_damage || 
				(damage == most_damage && def.effective_power()>groups[attacked].effective_power()) || 
				(damage == most_damage && def.effective_power()>groups[attacked].effective_power() && def.initiative > groups[attacked].initiative)  { // check ties order : target with the most effective power or initiative
				most_damage = damage
				attacked = def_id				
			}
		}
		fmt.Println("group",att_id+1,"of type",att.system,"would deal group",attacked+1,"with damage",most_damage)
		groups[attacked].attacked_by = att_id
		groups[attacked].attacks = attacked
	}

	// attack phase

	// sort by initiative, decreasing
	sort.Slice(groups, func(i,j int) bool {
		return groups[i].initiative > groups[j].initiative
		})
	for att_id,att := range groups {
		def := groups[att.attacks]
		units := att.damage(def) / def.hitpoints
		fmt.Println("Group",att_id,"attacks",att.attacks,"killed",units,"units")
	}
}

func main() {
	system := test_systems
	fmt.Println("- Immune to Infection")
	prepare_attack(system)
}
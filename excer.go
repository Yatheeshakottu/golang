package main
import(
	"fmt"
)
func main(){
	var n,m,games int
	MAX_players:=20000
	MIN_players_games:=1
	MAX_games:=30
	if n%2 !=0{
		break
	}
	if  n< MIN_players_games && n>MAX_players {
		if m< MIN_players_games && m>MAX_players{
			
		}
		break
	}
	mp:=map[string]int{
  teamone := "",
		teamtwo := "",

	}
	for:= range index  {
		int(1),int(n/2+1)
		mp= append(game[index-1])
		mp=append(game[int(index+n/2-1)])
	}
	for player in teamone{
		if mp(player)==None{
			mp[Player]=[]
			mp[player]+=teamtwo
		}
	}
	for player in teamtwo{
		if mp(player)==None{
			mp[player]=[]
			mp[player]+=teamone
		}
	}
	for player in mp{
		againstplayers:= mp[player]
againstplayers=list(set(againstplayers))
	}
	if player in againstplayers{
		againstplayers.remove(player)
		if len(againstplayers)!=n-1:{
			return false
		}
		return true
	}
	fmt.Print(check(2,1,[[1,2]]))
	fmt.Print(check(4,2,[[1,2,3,4],[4,3,1,2]]))
	fmt.Print(check(4,2[[1,2,3,4],[1,3,2,4]]))
	fmt.Print(check(6,6,[[1,6,3,4,5,2],[6,4,2,3,1,5],[4,2,1,5,6,3],[4,5,1,2,6,3],[3,2,5,1,6,4],[2,3,6,4,1,5]]))
	fmt.Print(check(6,6,[[3,1,4,5,6,2],[5,3,2,4,1,6],[5,3,6,4,2,1],[6,5,3,2,1,4],[5,4,1,2,6,3],[4,1,6,2,5,3]]))

}

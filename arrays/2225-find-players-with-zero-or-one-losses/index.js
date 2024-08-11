/**
 * @param {number[][]} matches
 * @return {number[][]}
 */
var findWinners = function(matches) {
  const winners = {}, losers = {}

  matches.forEach(([win, los] )=> {
    if(!winners.hasOwnProperty(win)){
      winners[win] = true
    }

    winners[los] = false

    if(!losers.hasOwnProperty(los)){
      losers[los] = true
    }else if(losers[los]){
      losers[los] = false
    }
  })

  const winRes = []
  const losRes = []

  Object.entries(winners).forEach(([key, value]) => {
    if(value) {
      winRes.push(parseInt(key))
    }
    if(losers[key]){
      losRes.push(parseInt(key))
    }

  })

  
  
  return [winRes, losRes]
};
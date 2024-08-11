
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum1 = function(nums, target) {
  for(let i = 0; i < nums.length; i++){
    for(let j = i+1; j < nums.length; j++){
        if(nums[i] + nums[j] === target){
            return [i, j]
        }
    }
  } 
};

console.time("2Fors")
let result1 = twoSum1([2,7,11,15], 9)
console.timeEnd('2Fors', result1)
console.log("Solution 1 (2 For's): ", result1)



// USING HASH TABLE
var twoSum2 = function(nums, target) {
  const map = new Map()
  
  for(let i = 0; i < nums.length; i++){
    let element =  nums[i]
    let value = map.get(element)
    if(value !== undefined){
      return [value, i]
    }

    map.set(target - element, i)
  } 
};

console.time('Map')
let result2 = twoSum2([2,7,11,15], 9)
console.timeEnd('Map', result1)
console.log("Solution 2 (Hash Table - Map): ", result2)




// USING HASH TABLE OBJECT
var twoSum3 = function(nums, target) {
  const object = {}
  
  for(let i = 0; i < nums.length; i++){
    let element =  nums[i]
    let value = object[element]
    if(value !== undefined){
      return [value, i]
    }

    object[target - element] = i 
  } 
};

console.time('object')
let result3 = twoSum3([2,7,11,15], 9)
console.timeEnd('object')
console.log("Solution 3 (Hash Table - Object):", result3)
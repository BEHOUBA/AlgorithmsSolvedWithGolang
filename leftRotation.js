// Challenge from HackerRank

// A left rotation operation on an array of size  "n"
// shifts each of the array's elements 1 unit to the left. 
// For example, if 2 left rotations are performed on array ,[1,2,3,4,5] 
// then the array would become [3,4,5,1,2].

// Given an array of  integers n and a number d, 
// perform n left rotations on the array. Then print the updated 
// array as a single line of space-separated integers.

function leftRotation(a, d) {
    // Complete this function
    let num1 = 0;
    for(let i = 0; i < d; i++) {
        num1 = a.splice(0, 1);
        a.splice(a.length, 0, num1[0]);
    }
    return a;
}


console.log(leftRotation([1, 2, 3, 4, 5], 4));
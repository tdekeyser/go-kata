## Some programming challenges

Here are some algorithmic problems; implement an **efficient** :rocket:, **maintainable** :wrench: and **tested** :rainbow: solution.

Don't forget to calculate the **runtime** for your algorithm!

#### Numbers

- **SumOfNumbers** Return the sum of each digit in the given number.

      For example, if the number is 259, then the sum should be 2+5+9 = 16.

- **kLargest** Write an efficient program for printing k largest elements in an array. Elements in array can be in any order.

      For example, if given array is [1, 23, 12, 9, 30, 2, 50] and you are asked for the largest 3 elements i.e.,
      k = 3 then your program should print 50, 30 and 23.

- **SumOfSubSequence** Given an array of integers, find the maximum sum of a subsequence of size k.
 
      Examples :
     
      Input : arr[] = {5, 5, 10, 100, 10, 5} , k = 3
      Output : 120
     
      Input : arr[] = {1, 2, 3} , k = 2
      Output : 5
     
      Input : arr[] = {1, 20, 3}, k = 1
      Output : 20

- **ListToNum** Given two numbers represented by two lists, write a function that returns sum list.
The sum list is list representation of addition of two input numbers.

      Example 1
        
      Input:
        First List: 5->6->3  // represents number 365
        Second List: 8->4->2 //  represents number 248
      Output
        Resultant list: 3->1->6  // represents number 613
        
      Example 2
        
      Input:
        First List: 7->5->9->4->6  // represents number 64957
        Second List: 8->4 //  represents number 48
      Output
        Resultant list: 5->0->0->5->6  // represents number 65005

- **PythagoreanTriplet** Given an array of integers, write a function that returns true if there is a triplet `(a, b, c)`
 that satisfies `a^2 + b^2 = c^2`.
                              
      Example:
              
      Input: arr[] = {3, 1, 4, 6, 5}
      Output: True
      There is a Pythagorean triplet (3, 4, 5).
              
      Input: arr[] = {10, 4, 6, 12, 5}
      Output: False
      There is no Pythagorean triplet.

- **NextGreaterElement** Given an array, print the Next Greater Element (NGE) for every element. The Next greater Element for an element x is the first
  greater element on the right side of x in array. Elements for which no greater element exist, consider next greater element as -1.
 
      Examples:
      a) For any array, rightmost element always has next greater element as -1.
      b) For an array which is sorted in decreasing order, all elements have next greater element as -1.
      c) For the input array [4, 5, 2, 25}, the next greater elements for each element are as follows.
     
      Element       NGE
         4      -->   5
         5      -->   25
         2      -->   25
         25     -->   -1
     
      d) For the input array [13, 7, 6, 12}, the next greater elements for each element are as follows.
     
        Element        NGE
         13      -->    -1
         7       -->     12
         6       -->     12
         12     -->     -1
        
- **MatrixRotation** Given an square matrix, turn it by 90 degrees in anti-clockwise direction without using any extra space.

      Examples :
    
      Input
       1  2  3
       4  5  6
       7  8  9
    
      Output:
       3  6  9
       2  5  8
       1  4  7
    
      Input:
       1  2  3  4
       5  6  7  8
       9 10 11 12
      13 14 15 16
    
      Output:
       4  8 12 16
       3  7 11 15
       2  6 10 14
       1  5  9 13

#### Trees and graphs

- **IceCreamFlavors** George goes to the ice cream shop which has n different ice cream flavors. He goes in with a list with m statements of the form
"I prefer flavor X to flavor Y". Your job is to determine if this list is consistent with itself. This means that if George
prefers X to Y and Y to Z, then he must prefer X to Z as well. This extends to having more than 3 flavors.

    Design an efficient algorithm for determining if George’s list is consistent with itself and state its run time.

- **CheapestRoads** Create an algorithm that finds the cheapest roads between N cities, based on the cost of building a cost between city A and B.

- **LowestCommonAncestor** Given a binary tree (not a binary search tree) and two values say n1 and n2, write a program to find the least common ancestor.

      Following is definition of LCA from Wikipedia:
        
      Let T be a rooted tree. The lowest common ancestor between two nodes 
      n1 and n2 is defined as the lowest node in T that has both n1 and n2 as descendants 
      (where we allow a node to be a descendant of itself).

#### Stacks, heaps, and queues

- **MinStack** Implement a stack; pop(), push(), peek(), and getMin(), all in O(1).
  Allowed to use Java Stack class.
  
  The difficulty is that when you pop, a new minimum should potentially be calculated; but it should be done in O(1).
  Space complexity doesn't matter.
  
  Hint for getMin() -> use a minStack to keep the smallest elements

#### Dynamic programming

- **StringPermutations** This method should return all permutations of a given string.
        
        For example, xy should result in the list xy,yx

- **PalindromePermutations** Given a string, write a function to check if it is a permutation of a palindrome. 
A palindrome is a word  or phrase that is the same forwards and  backwards.
A permutation is a rearrangement of letters. The palindrome does not need to be limited
to just dictionary words.

- **StockSpanProblem** The stock span problem is a financial problem where we have a series of n daily price quotes for a stock
  and we need to calculate span of stock’s price for all n days.
  The span Si of the stock’s price on a given day i is defined as the maximum number of consecutive days
  just before the given day, for which the price of the stock on the current day is less than or equal to its price on the given day.
  
        For example, if an array of 7 days prices is given as
            {100, 80, 60, 70, 60, 75, 85}, 
        then the span values for corresponding 7 days are
            {1, 1, 1, 2, 1, 4, 6}
            
- **Boggle** Given a dictionary, a method to do lookup in dictionary and a M x N board where every cell has one character.
  Find all possible words that can be formed by a sequence of adjacent characters. Note that we can move to any
  of 8 adjacent characters, but a word should not have multiple instances of same cell.
      
      Example:
      
      Input: dictionary[] = {"GEEKS", "FOR", "QUIZ", "GO"};
             boggle[][]   = {{'G', 'I', 'Z'},
                             {'U', 'E', 'K'},
                             {'Q', 'S', 'E'}};
            isWord(str): returns true if str is present in dictionary
                         else false.
      
      Output:  Following words of dictionary are present
               GEEKS
               QUIZ

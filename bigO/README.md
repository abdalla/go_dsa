# Big O

**Big O** notation is used in Computer Science to describe the performance or complexity of an algorithm. 

**Big O** is used to describe the execution time required or the space used (e.g. in memory or disk) by an algorithm.

To measure the efficiency of an algorithm, we need to consider two things:

1. **Time complexity** - How much time does it take to run completely?

2. **Space complexity** - How much extra space does it require in the process?
   
![bigO](../images/bigO_graph.png "bigO")
![bigO](../images/bigO_.png "bigO")

It doesn't depends on time by itself, it depends on how many steps we are performing

The main focus is to calculate the amount of work we do or the mumber of comparisons we perform.

## FIVE RULES TO SIMPLIFY Big O

1. Focus on Scalability

	* n tend to infinity

	* n ==> ♾️
	
2. Considering Worst Case Scenario

3. Remove all possible contants
	* O(6 + n) ==> O(n) 

4. Consider Different Variable for different inputs
	* O(n + m)

5. Remove all non-dominants
	* O(4 + n + 2nˆ2)  ==> [RULE 4] O(n + 2nˆ2) ==> [RULE 5] O(2nˆ2) ==> [RULE 4] O(nˆ2)
	* O(3nˆ2 + 3n + 500)  ==> O(nˆ2)
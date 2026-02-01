# 240. Search a 2D Matrix II
Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:

- Integers in each row are sorted in ascending from left to right.
- Integers in each column are sorted in ascending from top to bottom.

Constraints:

- m == matrix.length
- n == matrix[i].length
- 1 <= n, m <= 300
- -109 <= matrix[i][j] <= 109
- All the integers in each row are sorted in ascending order.
- All the integers in each column are sorted in ascending order.
- -109 <= target <= 109

# solution process
As the description suggests, there might be someway we can utilize this property of ascending order to achieve a running complexity less than O(m * n)

As we can easily spot that any number in the mastrix lies in the range of [M[topleft], [bottomright]], we can guess the invariant to be ``if the target falls into the range of between the top left and bottom right corner number, then it may exist. If not, the number definitly does not exists.``

So is it possible to remove a part of the range in the matrix, so we always remove a part that definitly does not contain the target and keep the rest?

There must be a loop in this solution, so when do we exit?

Is there any special situation that target falls into all ranges, how can we make the code terminate under such situation? 


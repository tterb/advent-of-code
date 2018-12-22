
def printRecipes(recipes, elfA, elfB):
    for n,i in enumerate(recipes):
        if n == elfA:
            print(f'({i})', end='')
        elif n == elfB:
            print(f'[{i}]', end='')
        else:
            print(f' {i} ', end='')
    print()

count = int(input())
# count = 9
recipeCount = count+10
elfA, elfB = 0, 1
recipes = [3,7]
while len(recipes) < recipeCount:
    temp = [int(d) for d in str(recipes[elfA]+recipes[elfB])]
    recipes.extend(temp)
    # printRecipes(recipes, elfA, elfB)
    elfA = (elfA+recipes[elfA]+1)%len(recipes)
    elfB = (elfB+recipes[elfB]+1)%len(recipes)
print(''.join(map(str, recipes[count:recipeCount])))
      
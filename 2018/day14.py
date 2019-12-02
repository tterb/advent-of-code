
def partOne(target):
    recipeCount = target+10
    elfA, elfB = 0, 1
    recipes = [3,7]
    while len(recipes) < recipeCount:
        temp = [int(d) for d in str(recipes[elfA]+recipes[elfB])]
        recipes.extend(temp)
        # printRecipes(recipes, elfA, elfB)
        elfA = (elfA+recipes[elfA]+1)%len(recipes)
        elfB = (elfB+recipes[elfB]+1)%len(recipes)
    print(''.join(map(str, recipes[target:recipeCount])))

def partTwo(target):
    n = len(str(target))
    elfA, elfB = 0, 1
    recipes = [3,7]
    val = ''
    found = False
    while not found:
        temp = [int(d) for d in str(recipes[elfA]+recipes[elfB])]
        for i in range(len(temp)):
            recipes.append(temp[i])
            if len(recipes) > n:
                val = ''.join(map(str, recipes[-n:]))
                if int(val) == target:
                    found = True
                    return len(recipes)-len(val)
            # printRecipes(recipes, elfA, elfB, val)
        elfA = (elfA+recipes[elfA]+1)%len(recipes)
        elfB = (elfB+recipes[elfB]+1)%len(recipes)

def printRecipes(recipes, elfA, elfB, val):
    for n,i in enumerate(recipes):
        if n == elfA:
            print(f'({i})', end='')
        elif n == elfB:
            print(f'[{i}]', end='')
        else:
            print(f' {i} ', end='')
    print(f'- {val}')

    
def main():
    target = int(input())
    # print(partOne(target))
    print(partTwo(target))
    

if __name__ == '__main__':
    main()

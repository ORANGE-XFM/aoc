# compiler for jumpcode
import ast 

tree = ast.parse("3 or 4 and (5 and not 6)").body[0].value
print(ast.dump(tree))

def eval_tree(t) : pass
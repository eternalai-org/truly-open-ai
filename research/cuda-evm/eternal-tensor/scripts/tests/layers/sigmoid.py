import numpy as np

np.set_printoptions(suppress=True)

a = np.random.rand(1, 10) * 10
# a = [
#         [0.80531654, 0.30585969, 0.13773904],
#         [0.06516666, 0.88074556, 0.03621177],
#         [0.3904885,  0.32787716, 0.29096427],
#     ]
print(a)

print(np.exp(a) / (np.exp(a)+1))


def softmax(x):
    exp_x = np.exp(x - np.max(x, axis=-1, keepdims=True))
    return exp_x / np.sum(exp_x, axis=-1, keepdims=True)

print(softmax(a))
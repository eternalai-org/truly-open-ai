import numpy as np

a = np.random.randint(0, 10, (2, 2))
print(a)

a_pad = np.pad(a, ((1, 2), (1, 3)), 'constant', constant_values=0)

print(a_pad)
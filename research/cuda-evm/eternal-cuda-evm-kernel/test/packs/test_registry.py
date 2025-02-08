from utils import log
import time
import enlighten
import itertools
from functools import wraps
import os

__test_fn_registry = []

'''
// example usage
    @wrap_test(
        name='abi random test',
        repeat=10,
        meta={
            'description': 'Test ABI encode/decode with random tensor'
        },
        params={
            spatial_size: [16, 32],
            channel_in: [1, 2, 4, 8, 512, 1024],
        }
    )
'''
def wrap_test(name, meta={}, repeat=1, params={}, checker=lambda **e: True):    
    def wrapper(fn):
        global __test_fn_registry
        __test_fn_registry.append((name, meta, fn, repeat, params, checker))
        @wraps(fn)
        def wrapped_fn(*args, **kwargs):
            return fn(*args, **kwargs)
        return wrapped_fn
    return wrapper

def repeat(fn, n, params_d = {}):
    for _ in range(n):
        yield fn(**params_d)

def run_tests(targets=[], **_):
    manager = enlighten.get_manager()
    pbar_global = manager.counter(total=len(__test_fn_registry), desc='Testing progress', unit='tests')
    status_format = '{program}{fill}{stage}{fill}{status}'

    pstatus = manager.status_bar(
        status_format=status_format,
        color='bold_slategray',
        program='Testing',
        stage='Preparing',
        status='Starting'
    )
    
    total_executed = 0
    
    os.makedirs('test_results', exist_ok=True)

    tests_results = []
    for i, (test_name, meta, fn, rp, params, checker) in enumerate(__test_fn_registry):
        if test_name not in targets:
            continue
        
        total_error, executed = 0, 0

        p_keys = list(params.keys()) 
        
        if len(p_keys) == 0:
            p_keys = ['_']
            p_combinations = [()]
        else:
            p_combinations = list(itertools.product(*[params[k] for k in p_keys]))
        
        p_combinations = [p for p in p_combinations if checker(**dict(zip(p_keys, p)))]
        pbar_local = manager.counter(total=rp * len(p_combinations), desc=f'Test {test_name}', unit='iterations')

        log(f'--- Running test {test_name} ---')

        for k, v in meta.items():
            log(f'   - {k}: {v}')
        
        test_results = []

        try:
            for p in p_combinations:
                params_d = dict(zip(p_keys, p))
                    
                
                log(f'Running test {test_name} with params:')
                for k, v in params_d.items():
                    log(f'  - {k}: {v}')

                for i, stats in enumerate(repeat(fn, rp, params_d)):

                    executed += 1
                    pbar_local.update()

                    stats['test_name'] = test_name
                    stats['params'] = params_d
                    
                    test_results.append(stats)
                    total_error += stats['error']

                    if i % 10 == 0:
                        pstatus.update(stage=f'Test {test_name}', status=f'Avg error: {total_error / executed:.8f} ({executed}/{rp * len(p_combinations)})')

        except KeyboardInterrupt:
            log('Interrupted')

        pbar_local.close()
        total_executed += executed

        with open(f'test_results/{test_name}.json', 'w') as f:
            import json
            json.dump(test_results, f)

        tests_results.append({
            'test_name': test_name,
            'meta': meta,
            'details': test_results,
            'total_error': total_error,
        })

        pbar_global.update()        
        
    pbar_global.close()
    return tests_results

def list_names():
    for i, (test_name, _, _, _, _, _) in enumerate(__test_fn_registry):
        yield test_name
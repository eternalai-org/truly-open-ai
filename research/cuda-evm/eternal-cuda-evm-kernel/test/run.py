from argparse import ArgumentParser 
from packs import run_tests, list_names
import os

def get_options():
    parser = ArgumentParser()    
    _list_names = list(list_names())
    parser.add_argument('--mode', type=str, default='test', choices=['test', 'benchmark'], help='Mode to run the tests', nargs='?')
    parser.add_argument('--target', type=str, default=_list_names if len(_list_names) > 0 else 'n/a', choices=_list_names, help='Name of the test to run', nargs='+')
    parser.add_argument('-e', '--accepted-error', type=float, default=1e-6, help='Accepted error for the tests')
    return parser.parse_args()

def main():
    options = get_options()
    print('Starting tests...')

    results = run_tests(
        targets=options.target, 
        benchmark=options.mode == 'benchmark', 
        accepted_error=options.accepted_error
    )

    print('Tests finished.')
    print(results)
        
if __name__ == '__main__':
    main()
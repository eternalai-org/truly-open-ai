import os
import sys
import argparse
import requests
import numpy as np
from eai.version import __version__
from eai.utils import Logger, ENV_PATH, ETHER_PER_WEI
from eai.utils import create_web3_account, publisher
from eai.network_config import NETWORK
from eai.func import transform, check, get_model, transfer_model

DEFAULT_NETWORK = "mainnet"
DEFAULT_DUMP_FILE_NAME = "secret"


def parse_args():
    parser = argparse.ArgumentParser(
        description="Tool for managing and deploying machine learning models on-chain.")

    subparsers = parser.add_subparsers(
        dest='command', help="Commands to execute for EternalAI.")
    subparsers.add_parser('version', help="Show the version of the toolkit.")
    subparsers.add_parser(
        'upgrade', help="Upgrade the toolkit to the latest version.")
    # Wallet command
    parser_wallet = subparsers.add_parser(
        'wallet', help="Commands to execute for wallet operations.")
    wallet_commands = parser_wallet.add_subparsers(
        dest='subcommand', help="Subcommands to execute for wallet.")

    wallet_create = wallet_commands.add_parser(
        'create', help='Create a new wallet')
    wallet_create.add_argument(
        '--name',
        '-na',
        action='store',
        default=DEFAULT_DUMP_FILE_NAME,
        type=str,
        help=f"Name of the dump file. Default is {DEFAULT_DUMP_FILE_NAME}."
    )
    wallet_create.add_argument(
        "--network",
        "-ne",
        action='store',
        default=None,
        type=str,
        help=f"Network mode. Specify the network configuration. Default is {DEFAULT_NETWORK}."
    )
    wallet_import = wallet_commands.add_parser(
        'import', help='Import a wallet from private key.')
    wallet_import.add_argument(
        "--private-key",
        "-p",
        action='store',
        type=str,
        help="Private key for the wallet."
    )
    wallet_import.add_argument(
        "--file",
        "-f",
        action='store',
        type=str,
        help="Path to load private key from file."
    )
    wallet_import.add_argument(
        "--network",
        "-ne",
        action='store',
        default=None,
        type=str,
        help=f"Network mode. Specify the network configuration. Default is {DEFAULT_NETWORK}."
    )
    wallet_balance = wallet_commands.add_parser(
        'balance', help='Check the balance of the wallet.')
    wallet_transactions = wallet_commands.add_parser(
        'transactions', help='List all transactions for the wallet.')
    wallet_transactions.add_argument(
        "--output-path",
        "-o",
        action='store',
        default=None,
        type=str,
        help="Output path for saving transactions list."
    )
    wallet_transactions.add_argument(
        "--network",
        "-ne",
        action='store',
        default=None,
        type=str,
        help=f"Network mode. Specify the network configuration. Default is {DEFAULT_NETWORK}."
    )
    wallet_faucet = wallet_commands.add_parser(
        'faucet', help='Request EAI from the faucet.')
    wallet_faucet.add_argument(
        "--network",
        "-ne",
        action='store',
        default=None,
        type=str,
        help=f"Network mode. Specify the network configuration. Default is {DEFAULT_NETWORK}."
    )
    wallet_receive = wallet_commands.add_parser(
        'receive', help='Address to receive EAI tokens.')
    wallet_send = wallet_commands.add_parser(
        'send', help='Send EAI tokens to an address.')
    wallet_send.add_argument(
        "--recipient",
        "-r",
        action='store',
        type=str,
        help="Address of the recipient."
    )
    wallet_send.add_argument(
        "--eternal-id",
        "-id",
        action='store',
        type=str,
        help="EternalAI model ID."
    )
    wallet_send.add_argument(
        "--eternal-address",
        "-a",
        action='store',
        type=str,
        help="EternalAI model address."
    )
    wallet_send.add_argument(
        "--network",
        "-ne",
        action='store',
        default=None,
        type=str,
        help=f"Network mode. Specify the network configuration. Default is {DEFAULT_NETWORK}."
    )
    wallet_deposit = wallet_commands.add_parser(
        'deposit', help='Generate deposit address.')
    wallet_deposit.add_argument(
        "--output-path",
        "-o",
        action='store',
        default=None,
        type=str,
        help="Output path for saving deposit metadata."
    )
    wallet_deposit.add_argument(
        "--network",
        "-ne",
        action='store',
        default=None,
        type=str,
        help="Network mode. Specify the network configuration."
    )

    # Eternal command
    parser_eternal = subparsers.add_parser(
        'eternal', help='Commands to execute for eternal operations.')
    eternal_commands = parser_eternal.add_subparsers(
        dest='subcommand', help="Subcommands to execute for eternal.")

    eternal_transform = eternal_commands.add_parser(
        'transform', help='Transform a Keras model to the EternalAI chain.')
    eternal_transform.add_argument(
        "--format",
        "-fo",
        action='store',
        type=str,
        default='keras3',
        choices=['keras2', 'keras3'],
        help="Format of the model to deploy on-chain. Default is 'keras3'. Choices: 'keras2', 'keras3'."
    )
    eternal_transform.add_argument(
        "--file",
        "-f",
        action='store',
        default=None,
        type=str,
        help="Path to the model file."
    )
    eternal_transform.add_argument(
        "--url",
        "-u",
        action='store',
        default=None,
        type=str,
        help="URL to the model file."
    )
    eternal_transform.add_argument(
        "--name",
        "-na",
        action='store',
        type=str,
        default="Unnamed Model",
        help="Name of the model. Default is 'Unnamed Model'."
    )
    eternal_transform.add_argument(
        "--network",
        "-ne",
        action='store',
        type=str,
        default=None,
        help="Network mode. Specify the network configuration."
    )
    eternal_transform.add_argument(
        "--output-path",
        "-o",
        action='store',
        default=None,
        type=str,
        help="Output path for saving transformed model metadata."
    )

    eternal_list = eternal_commands.add_parser(
        'list', help='List all deployed models')
    eternal_list.add_argument(
        "--network",
        "-ne",
        action='store',
        default=None,
        type=str,
        help="Network mode. Specify the network configuration."
    )
    eternal_list.add_argument(
        "--output-path",
        "-o",
        action='store',
        default=None,
        type=str,
        help="Output path for saving deployed models list."
    )
    eternal_check = eternal_commands.add_parser(
        'check', help='Check the model format.')
    eternal_check.add_argument(
        "--format",
        "-fo",
        action='store',
        type=str,
        default='keras3',
        choices=['keras2', 'keras3'],
        help="Format of the model to deploy on-chain. Default is 'keras3'. Choices: 'keras2', 'keras3'."
    )
    eternal_check.add_argument(
        "--file",
        "-f",
        action='store',
        default=None,
        type=str,
        help="Path to the model file."
    )
    eternal_check.add_argument(
        "--url",
        "-u",
        action='store',
        default=None,
        type=str,
        help="URL to the model file."
    )
    eternal_check.add_argument(
        "--output-path",
        "-o",
        action='store',
        default=None,
        type=str,
        help="Output path for saving check result."
    )
    # Call command
    parser_call = subparsers.add_parser(
        'call', help='Commands to execute for call operations.')
    call_commands = parser_call.add_subparsers(
        dest='subcommand', help="Subcommands to execute for call.")

    call_predict = call_commands.add_parser(
        'predict', help="Call the predict function of a deployed model.")
    call_predict.add_argument(
        "--input",
        "-i",
        action='store',
        type=str,
        help="Input data for 'eai call predict' command."
    )
    call_predict.add_argument(
        "--eternal-address",
        "-e",
        action='store',
        type=str,
        help="Address of the model for 'eai call' command."
    )
    call_predict.add_argument(
        "--output-path",
        "-o",
        action='store',
        default=None,
        type=str,
        help="Output path for saving prediction result."
    )
    return parser.parse_known_args()


def version_command():
    Logger.success(f"✨ EternalAI Toolkit - Version: {__version__} ✨")


def upgrade_command():
    os.system(
        "pip install --upgrade git+https://github.com/eternalai-org/eternalai.git")


def wallet_create(**kwargs):
    if "PRIVATE_KEY" in os.environ:
        Logger.warning(
            "Wallet already set. This will overwrite the existing wallet in .env file."
        )
    account = create_web3_account()
    network = kwargs['network'] if kwargs['network'] is not None else DEFAULT_NETWORK
    Logger.success(f"Private key: {account['private_key']}")
    Logger.success(f"Address: {account['address']}")
    env_config = {
        "PRIVATE_KEY": account['private_key'],
        "NETWORK_MODE": network
    }
    with open(ENV_PATH, "w") as f:
        for key, value in env_config.items():
            f.write(f"{key}={value}\n")
            os.environ[key] = str(value)
    if not os.path.exists(network):
        os.makedirs(network)

    file_path = os.path.join(network, kwargs['name'] + ".txt")
    with open(file_path, "w") as f:
        f.write(f"{account['private_key']}")
    Logger.success(
        f"Wallet created and set successfully. Private key written to {file_path}.")


def wallet_import(**kwargs):
    if "PRIVATE_KEY" in os.environ:
        Logger.warning(
            "Wallet already set. This will overwrite the existing wallet in .env file."
        )
    if kwargs['private-key'] is None:
        if kwargs['file'] is None:
            Logger.error(
                "Please provide either --private-key or --file for importing wallet.")
            sys.exit(2)
        else:
            with open(kwargs['file'], "r") as f:
                private_key = f.read()
            kwargs['private-key'] = private_key
    env_config = {
        "PRIVATE_KEY": kwargs['private-key'],
        "NETWORK_MODE": kwargs['network'] if kwargs['network'] is not None else DEFAULT_NETWORK
    }
    with open(ENV_PATH, "w") as f:
        for key, value in env_config.items():
            f.write(f"{key}={value}\n")
            os.environ[key] = str(value)
    Logger.success("Wallet imported and set successfully.")


def wallet_balance():
    if "PRIVATE_KEY" not in os.environ:
        Logger.error(
            "Wallet not found. Please use command 'eai wallet create' or 'eai wallet import' to create a new wallet or restore from private key.")
        sys.exit(2)
    balance = {
        "testnet": 0,
        "mainnet": 0
    }
    tesnet_explorer = NETWORK["testnet"]["EXPLORER_ENDPOINT"]
    testnet_url = f"{tesnet_explorer}?module=account&action=balancemulti&address={publisher()}"
    response = requests.get(testnet_url)
    if response.status_code == 200:
        data = response.json()
        result = data.get("result", None)
        if result is not None:
            balance["testnet"] = result[0]["balance"]
    balance["testnet"] = float(balance["testnet"]) / ETHER_PER_WEI
    mainnet_explorer = NETWORK["mainnet"]["EXPLORER_ENDPOINT"]
    testnet_url = f"{mainnet_explorer}?module=account&action=balancemulti&address={publisher()}"
    response = requests.get(testnet_url).json()
    if response["message"] == "OK":
        result = response.get("result", None)
        if result is not None:
            balance["mainnet"] = result[0]["balance"]
        balance["mainnet"] = float(balance["mainnet"]) / ETHER_PER_WEI
        Logger.success(f"Testnet: {balance['testnet']}")
        Logger.success(f"Mainnet: {balance['mainnet']}")
    else:
        Logger.error("Failed to get balance.")
    return balance


def wallet_transactions(**kwargs):
    if "PRIVATE_KEY" not in os.environ:
        Logger.error(
            "Wallet not found. Please use command 'eai wallet create' or 'eai wallet import' to create a new wallet or restore from private key.")
        sys.exit(2)
    network = kwargs["network"] if kwargs["network"] is not None else os.environ["NETWORK_MODE"]
    transactions = []
    explorer_endpoint = NETWORK[network]["EXPLORER_ENDPOINT"]
    url = f"{explorer_endpoint}?module=account&action=txlist&address={publisher()}&sort=desc"
    response = requests.get(url)
    if response.status_code == 200:
        data = response.json()
        transactions = data.get("result", [])
    for idx, transaction in enumerate(transactions):
        tx_hash = transaction["hash"]
        timeStamp = transaction["timeStamp"]
        Logger.success(f"tx: {tx_hash}, timestamp: {timeStamp}.")
    if kwargs["output-path"] is not None:
        with open(kwargs["output-path"], "w") as f:
            for idx, transaction in enumerate(transactions):
                tx_hash = transaction["hash"]
                timeStamp = transaction["timeStamp"]
                f.write(f"tx: {tx_hash}, timestamp: {timeStamp}.\n")
        Logger.success(f"Transactions written to {kwargs['output-path']}.")


def wallet_faucet(**kwargs):
    if "PRIVATE_KEY" not in os.environ:
        Logger.error(
            "Wallet not found. Please use command 'eai wallet create' or 'eai wallet import' to create a new wallet or restore from private key.")
        sys.exit(2)
    network = kwargs["network"] if kwargs["network"] is not None else os.environ["NETWORK_MODE"]
    if network == "testnet":
        faucet_endpoint = NETWORK[network]["FAUCET_ENDPOINT"]
        url = f"{faucet_endpoint}/{publisher()}"
        response = requests.post(url).json()
        if response["status"] == 1:
            Logger.success("Faucet request successful.")
        else:
            error = response.get("error", "Something went wrong.")
            Logger.error(f"Faucet request failed: {error}")
    else:
        Logger.error("'eai wallet faucet' is only available on testnet.")
        sys.exit(2)


def wallet_receive():
    if "PRIVATE_KEY" not in os.environ:
        Logger.error(
            "Wallet not found. Please use command 'eai wallet create' or 'eai wallet import' to create a new wallet or restore from private key.")
        sys.exit(2)
    address = publisher()
    Logger.success(f"Address: {address}")
    return address


def wallet_send(**kwargs):
    if "PRIVATE_KEY" not in os.environ:
        Logger.error(
            "Wallet not found. Please use command 'eai wallet create' or 'eai wallet import' to create a new wallet or restore from private key.")
        sys.exit(2)
    if kwargs["recipient"] is None:
        Logger.error(
            "Recipient address must be provided for 'eai wallet send' command.")
        sys.exit(2)
    if kwargs["eternal-id"] is None:
        if kwargs["eternal-address"] is None:
            Logger.error(
                "Either EternalAI model ID or address must be provided for 'eai wallet send' command.")
            sys.exit(2)
        else:
            address = kwargs["eternal-address"]
            endpoint = NETWORK[os.environ["NETWORK_MODE"]
                               ]["MODEL_INFO_BY_ADDRESS"]
            url = f"{endpoint}/{address}"
            response = requests.get(url)
            response = response.json()
            if response["status"] == 1:
                kwargs["eternal-id"] = response["data"]["model_id"]
    transfer_model(kwargs["eternal-id"],
                   kwargs["recipient"], kwargs["network"])


def wallet_deposit(**kwargs):
    if "PRIVATE_KEY" not in os.environ:
        Logger.error(
            "Wallet not found. Please use command 'eai wallet create' or 'eai wallet import' to create a new wallet or restore from private key.")
        sys.exit(2)
    network = kwargs["network"] if kwargs["network"] is not None else os.environ["NETWORK_MODE"]
    output_path = kwargs["output-path"]
    if network == "testnet":
        Logger.error("'eai wallet deposit' is not available on testnet.")
        sys.exit(2)
    deposit_endpoint = NETWORK[network]["DEPOSIT_ENDPOINT"]
    payload = {
        "tcAddress": publisher(),
        "tcTokenID": "0x0000000000000000000000000000000000000000",
        "toChainID": 43338,
        "fromChainID": 1
    }
    headers = {
        "Content-Type": "application/json"
    }
    response = requests.post(
        deposit_endpoint, json=payload, headers=headers).json()
    if response["status"]:
        deposit_address = response["data"]["depositAddress"]
        Logger.success(f"Deposit address: {deposit_address}")
        if output_path is not None:
            with open(output_path, "w") as f:
                f.write(f"Deposit address: {deposit_address}")
            Logger.success(f"Deposit address written to {output_path}.")


def handle_wallet(args):
    if args.subcommand == "create":
        kwargs = {
            "name": args.name,
            "network": args.network
        }
        wallet_create(**kwargs)
    elif args.subcommand == "import":
        kwargs = {
            "private-key": args.private_key,
            "network": args.network,
            'file': args.file
        }
        wallet_import(**kwargs)
    elif args.subcommand == "transactions":
        kwargs = {
            "output-path": args.output_path,
            "network": args.network
        }
        wallet_transactions(**kwargs)
    elif args.subcommand == "balance":
        wallet_balance()
    elif args.subcommand == "faucet":
        kwargs = {
            "network": args.network
        }
        wallet_faucet(**kwargs)
    elif args.subcommand == "receive":
        wallet_receive()
    elif args.subcommand == "send":
        kwargs = {
            "recipient": args.recipient,
            "eternal-id": args.eternal_id,
            "eternal-address": args.eternal_address,
            "network": args.network
        }
        wallet_send(**kwargs)
    elif args.subcommand == "deposit":
        kwargs = {
            "output-path": args.output_path,
            "network": args.network
        }
        wallet_deposit(**kwargs)
    else:
        Logger.error(
            f"Subcommand {args.subcommand} not work for eai {args.command}.")
        sys.exit(2)


def eternal_transform(**kwargs):
    if "PRIVATE_KEY" not in os.environ:
        Logger.error(
            "Wallet not found. Please use command 'eai wallet create' or 'eai wallet import' to create a new wallet or restore from private key.")
        sys.exit(2)
    if kwargs['file'] is None:
        if kwargs['url'] is None:
            Logger.error(
                'Please provide either --file or --url for the model.')
            sys.exit(2)
        else:
            Logger.info(f"Downloading model from {kwargs['url']} ...")
            import tempfile
            with tempfile.NamedTemporaryFile(delete=True) as temp_file:
                temp_file_name = temp_file.name + ".keras"
            os.system(f"wget {kwargs['url']} -O {temp_file_name}")
            kwargs['file'] = temp_file_name
    eternal = transform(kwargs['file'], kwargs['name'],
                        kwargs['format'], network_mode=kwargs['network'])
    eternal.to_json(kwargs['output-path'])


def eternal_check(**kwargs):
    if kwargs['file'] is None:
        if kwargs['url'] is None:
            Logger.error(
                'Please provide either --file or --url for the model.')
            sys.exit(2)
        else:
            Logger.info(f"Downloading model from {kwargs['url']} ...")
            import tempfile
            with tempfile.NamedTemporaryFile(delete=True) as temp_file:
                temp_file_name = temp_file.name + ".keras"
            os.system(f"wget {kwargs['url']} -O {temp_file_name}")
            kwargs['file'] = temp_file_name
    check(kwargs['file'], kwargs['format'], kwargs['output-path'])


def eternal_list(**kwargs):
    if "PRIVATE_KEY" not in os.environ:
        Logger.error(
            "Wallet not found. Please use command 'eai wallet create' or 'eai wallet import' to create a new wallet or restore from private key.")
        sys.exit(2)
    network = kwargs["network"] if kwargs["network"] is not None else os.environ["NETWORK_MODE"]
    list_model_endpoint = NETWORK[network]["LIST_MODEL_ENDPOINT"]
    url = f"{list_model_endpoint}?address={publisher()}&public_only=false&limit=20&offset=0"
    response = requests.get(url).json()
    deployed_models = []
    if response["status"] == 1:
        models = response["data"]
        for model_info in models:
            if model_info["status"] == "deployed":
                name = model_info["model_name"]
                address = model_info["model_address"]
                id = model_info["model_id"]
                price = 0
                owner = model_info["owner"]["address"]
                model_info_str = f"Model: {name}, Address: {address}, ID: {id}, Price: {price}, Owner: {owner}"
                Logger.success(model_info_str)
                deployed_models.append(model_info_str)
        if kwargs["output-path"] is not None:
            with open(kwargs["output-path"], "w") as f:
                for model_info_str in deployed_models:
                    f.write(f"{model_info_str}\n")
    else:
        Logger.error(
            f"Failed to get deployed models from EternalAI's {network}.")


def handle_eternal(args):
    if args.subcommand == "transform":
        kwargs = {
            "file": args.file,
            "url": args.url,
            "name": args.name,
            "format": args.format,
            "network": args.network,
            "output-path": args.output_path
        }
        eternal_transform(**kwargs)
    elif args.subcommand == "check":
        kwargs = {
            "file": args.file,
            "url": args.url,
            "format": args.format,
            "output-path": args.output_path
        }
        eternal_check(**kwargs)
    elif args.subcommand == "list":
        kwargs = {
            "network": args.network,
            "output-path": args.output_path
        }
        eternal_list(**kwargs)
    else:
        Logger.error(
            f"Subcommand {args.subcommand} not work for eai {args.command}.")
        sys.exit(2)


def call_predict(**kwargs):
    model = get_model(kwargs['eternal-address'])
    input_data = np.load(kwargs['input'])
    output_arr = model.predict([input_data], kwargs['output-path'])
    return output_arr


def handle_call(args):
    if args.subcommand == "predict":
        kwargs = {
            "input": args.input,
            "eternal-address": args.eternal_address,
            "output-path": args.output_path
        }
        call_predict(**kwargs)
    else:
        Logger.error(
            f"Subcommand {args.subcommand} not work for eai {args.command}.")
        sys.exit(2)


@Logger.catch
def main():
    known_args, unknown_args = parse_args()
    for arg in unknown_args:
        Logger.error(f'unknown command or argument: {arg}')
        sys.exit(2)

    if known_args.command == "version":
        version_command()
    if known_args.command == "upgrade":
        upgrade_command()
    elif known_args.command == "wallet":
        handle_wallet(known_args)
    elif known_args.command == "eternal":
        handle_eternal(known_args)
    elif known_args.command == "call":
        handle_call(known_args)


if (__name__ == "__main__"):
    main()

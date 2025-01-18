import logging
import time
from .telegram_kit import send_message

# Configure basic logging (can be customized per project)
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
)


def log_function_call(func):
    """
    Logs the name of the function being called, its arguments, and return value
    """

    def wrapper(*args, **kwargs):
        logging.info(
            f"Function `{func.__name__}` called with args: {args} and kwargs: {kwargs}"
        )
        result = func(*args, **kwargs)
        logging.info(f"Function `{func.__name__}` returned: {result}")
        return result

    return wrapper


def log_execution_time(func):
    """
    Logs the execution time of the function.
    """

    def wrapper(*args, **kwargs):
        start_time = time.time()
        result = func(*args, **kwargs)
        elapsed_time = time.time() - start_time
        logging.info(
            f"Function `{func.__name__}` executed in {elapsed_time:.4f} seconds"
        )
        return result

    return wrapper

import traceback

def log_on_error(func):
    """
    Logs an error if the function raises an exception.
    """

    def wrapper(*args, **kwargs):
        try:
            return func(*args, **kwargs)
        except Exception as e:
            logging.error(
                f"Function `{func.__name__}` raised an error: {e}", exc_info=True
            )
            
            raise

    return wrapper


def log_on_error_and_raise_alert(alert_room: str = None):
    """
    Logs an error if the function raises an exception.
    """

    def decorator(func):
        def wrapper(*args, **kwargs):
            try:
                return func(*args, **kwargs)
            except Exception as e:
                logging.error(
                    f"Function `{func.__name__}` raised an error: {e}", exc_info=True
                )

                msg = '## Function `{}` raised an error: {}; \n\n## Inputs: \n-args: {}\n-kwargs {} \n\n## Traceback: \n\n ```bash\n{}\n```'.format(
                    func.__name__, e, 
                    args, kwargs, traceback.format_exc()
                )

                alert_room is None or send_message('junk_notifications', msg, room=alert_room)
                raise e
        return wrapper
    return decorator


def log_custom_message(message: str, level=logging.INFO):
    """
    Logs a custom message before calling the function.
    """

    def decorator(func):
        def wrapper(*args, **kwargs):
            logging.log(level, f"{message} - Function `{func.__name__}` is starting.")
            result = func(*args, **kwargs)
            logging.log(level, f"{message} - Function `{func.__name__}` completed.")
            return result

        return wrapper

    return decorator


# Combine multiple decorators (example)
def log_all(func):
    """
    Combines logging of function calls, execution time, and error handling.
    """

    @log_function_call
    @log_execution_time
    @log_on_error
    def wrapper(*args, **kwargs):
        return func(*args, **kwargs)

    return wrapper

from pathlib import Path
import yaml
from logging import Logger
import sys

CONFIG_FILE_NAME = "config.yml"
DATA_DIR_PATH = 'data_dir'


class ConfigLoader:

    def __init__(self, logger: Logger):
        self._logger = logger

    def load_config(self):
        source_path = Path(__file__).resolve()
        config_file_path = source_path.parent / CONFIG_FILE_NAME
        stream = open(config_file_path, 'r')
        config = yaml.safe_load(stream)
        data_dir_path = self._extract_data_dir(config)
        return {
            DATA_DIR_PATH: data_dir_path
        }

    def _extract_data_dir(self, config) -> Path:
        try:
            path_str = config[DATA_DIR_PATH]
            data_dir_path = Path(path_str)
            if not data_dir_path.is_dir():
                raise FileNotFoundError(data_dir_path)
            return data_dir_path
        except KeyError as e:
            self._logger.fatal(
                f"Mandatory Configuration property '{e.args[0]}' is not defined")
            sys.exit(1)
        except FileNotFoundError as e:
            self._logger.fatal(
                f"The data directory '{e.args[0]}' does not exist")
        except:
            self._logger.fatal(f"Unknown error reading configuration file")
            sys.exit(1)

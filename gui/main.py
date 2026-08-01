import sys

from PySide6.QtWidgets import QApplication

from main_window import ChimeraTaskWindow


def main():
    app = QApplication(sys.argv)

    window = ChimeraTaskWindow()
    window.show()

    sys.exit(app.exec())


if __name__ == "__main__":
    main()

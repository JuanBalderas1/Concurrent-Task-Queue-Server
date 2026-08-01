CHIMERA_STYLE = """
QMainWindow {
    background-color: #202124;
}

QWidget {
    color: #E4E4E4;
    font-family: "Segoe UI", Arial, sans-serif;
    font-size: 14px;
}

QFrame#headerPanel,
QFrame#memberPanel,
QFrame#taskPanel,
QFrame#composerPanel {
    background-color: #2B2D31;
    border: 1px solid #111214;
    border-radius: 10px;
}

QLabel#appTitle {
    color: #F28C28;
    font-size: 26px;
    font-weight: 700;
}

QLabel#subtitle {
    color: #A7A7A7;
    font-size: 13px;
}

QLabel#sectionTitle {
    color: #F28C28;
    font-size: 18px;
    font-weight: 700;
}

QLabel#serverOnline {
    color: #58A66A;
    background-color: #24372A;
    border: 1px solid #3F774B;
    border-radius: 10px;
    padding: 5px 10px;
    font-weight: 700;
}

QLabel#serverOffline {
    color: #E16B6B;
    background-color: #412626;
    border: 1px solid #7B3F3F;
    border-radius: 10px;
    padding: 5px 10px;
    font-weight: 700;
}

QLineEdit,
QComboBox,
QTextEdit {
    background-color: #17191C;
    color: #E4E4E4;
    border: 1px solid #111214;
    border-radius: 6px;
    padding: 7px;
    selection-background-color: #F28C28;
    selection-color: #17191C;
}

QLineEdit:focus,
QComboBox:focus,
QTextEdit:focus {
    border: 1px solid #F28C28;
}

QPushButton {
    min-height: 34px;
    background-color: #35383D;
    color: #E4E4E4;
    border: 1px solid #111214;
    border-radius: 6px;
    padding: 4px 13px;
    font-weight: 600;
}

QPushButton:hover {
    background-color: #41454B;
    border-color: #F28C28;
}

QPushButton#primaryButton {
    background-color: #F28C28;
    color: #17191C;
    border-color: #F28C28;
    font-weight: 700;
}

QPushButton#primaryButton:hover {
    background-color: #FFAA4C;
    border-color: #FFAA4C;
}

QListWidget {
    background-color: #17191C;
    color: #E4E4E4;
    border: 1px solid #111214;
    border-radius: 6px;
    padding: 5px;
    outline: none;
}

QListWidget::item {
    border-radius: 5px;
    padding: 8px;
    margin: 2px;
}

QListWidget::item:hover {
    background-color: #35383D;
}

QListWidget::item:selected {
    background-color: #F28C28;
    color: #17191C;
}

QStatusBar {
    background-color: #17191C;
    color: #A7A7A7;
    border-top: 1px solid #111214;
}

QStatusBar::item {
    border: none;
}
"""

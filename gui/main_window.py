from PySide6.QtCore import QTimer, Qt
from PySide6.QtWidgets import (
    QComboBox,
    QDialog,
    QDialogButtonBox,
    QFormLayout,
    QFrame,
    QHBoxLayout,
    QLabel,
    QLineEdit,
    QListWidget,
    QListWidgetItem,
    QMainWindow,
    QPushButton,
    QTextEdit,
    QVBoxLayout,
    QWidget,
)

from api_client import APIError, ChimeraAPIClient
from styles import CHIMERA_STYLE


class AddMemberDialog(QDialog):
    def __init__(self, parent=None):
        super().__init__(parent)

        self.setWindowTitle("Add Chimera Member")
        self.setMinimumWidth(380)

        self.name_input = QLineEdit()
        self.contact_input = QLineEdit()
        self.role_input = QLineEdit()

        self.name_input.setPlaceholderText(
            "Example: Worker Gamma"
        )
        self.contact_input.setPlaceholderText(
            "Example: gamma@chimera.local"
        )
        self.role_input.setPlaceholderText(
            "Example: Worker"
        )

        form_layout = QFormLayout()
        form_layout.addRow(
            "Name:",
            self.name_input,
        )
        form_layout.addRow(
            "Contact:",
            self.contact_input,
        )
        form_layout.addRow(
            "Role:",
            self.role_input,
        )

        buttons = QDialogButtonBox(
            QDialogButtonBox.Save
            | QDialogButtonBox.Cancel
        )

        buttons.accepted.connect(self.accept)
        buttons.rejected.connect(self.reject)

        layout = QVBoxLayout(self)
        layout.addLayout(form_layout)
        layout.addWidget(buttons)

    def member_data(self):
        return {
            "name": self.name_input.text().strip(),
            "contact": self.contact_input.text().strip(),
            "role": self.role_input.text().strip(),
        }


class ChimeraTaskWindow(QMainWindow):
    def __init__(self):
        super().__init__()

        self.setWindowTitle("Chimera Task Server")
        self.resize(1250, 780)
        self.setMinimumSize(1000, 650)
        self.setStyleSheet(CHIMERA_STYLE)

        self.api = ChimeraAPIClient()
        self.members = []

        self.build_interface()
        self.connect_signals()

        self.refresh_all()

        self.refresh_timer = QTimer(self)
        self.refresh_timer.setInterval(2000)
        self.refresh_timer.timeout.connect(
            self.refresh_tasks
        )
        self.refresh_timer.start()

    # ------------------------------------------------------------
    # Interface
    # ------------------------------------------------------------

    def build_interface(self):
        central_widget = QWidget()
        central_layout = QVBoxLayout(central_widget)

        central_layout.setContentsMargins(
            18,
            18,
            18,
            12,
        )
        central_layout.setSpacing(14)

        central_layout.addWidget(
            self.create_header()
        )

        content_layout = QHBoxLayout()
        content_layout.setSpacing(14)

        content_layout.addWidget(
            self.create_member_panel(),
            1,
        )

        content_layout.addWidget(
            self.create_task_panel(),
            3,
        )

        central_layout.addLayout(
            content_layout,
            1,
        )

        central_layout.addWidget(
            self.create_composer_panel()
        )

        self.setCentralWidget(central_widget)

    def create_header(self):
        header = QFrame()
        header.setObjectName("headerPanel")

        layout = QHBoxLayout(header)
        layout.setContentsMargins(18, 14, 18, 14)

        title_layout = QVBoxLayout()
        title_layout.setSpacing(2)

        title = QLabel("CHIMERA TASK SERVER")
        title.setObjectName("appTitle")

        subtitle = QLabel(
            "Queue, route, and monitor work "
            "between connected members."
        )
        subtitle.setObjectName("subtitle")

        title_layout.addWidget(title)
        title_layout.addWidget(subtitle)

        self.server_status = QLabel(
            "● CHECKING SERVER"
        )
        self.server_status.setAlignment(
            Qt.AlignCenter
        )

        self.refresh_button = QPushButton(
            "REFRESH"
        )

        layout.addLayout(title_layout)
        layout.addStretch()
        layout.addWidget(self.server_status)
        layout.addWidget(self.refresh_button)

        return header

    def create_member_panel(self):
        panel = QFrame()
        panel.setObjectName("memberPanel")
        panel.setMinimumWidth(250)

        layout = QVBoxLayout(panel)
        layout.setContentsMargins(14, 14, 14, 14)
        layout.setSpacing(10)

        title = QLabel("MEMBERS")
        title.setObjectName("sectionTitle")

        self.member_list = QListWidget()

        self.add_member_button = QPushButton(
            "ADD MEMBER"
        )
        self.add_member_button.setObjectName(
            "primaryButton"
        )

        layout.addWidget(title)
        layout.addWidget(self.member_list, 1)
        layout.addWidget(self.add_member_button)

        return panel

    def create_task_panel(self):
        panel = QFrame()
        panel.setObjectName("taskPanel")

        layout = QVBoxLayout(panel)
        layout.setContentsMargins(14, 14, 14, 14)
        layout.setSpacing(10)

        title_row = QHBoxLayout()

        title = QLabel("TASK CONVERSATION")
        title.setObjectName("sectionTitle")

        self.task_count_label = QLabel(
            "0 TASKS"
        )
        self.task_count_label.setObjectName(
            "subtitle"
        )

        title_row.addWidget(title)
        title_row.addStretch()
        title_row.addWidget(
            self.task_count_label
        )

        self.task_list = QListWidget()

        layout.addLayout(title_row)
        layout.addWidget(self.task_list, 1)

        return panel

    def create_composer_panel(self):
        panel = QFrame()
        panel.setObjectName("composerPanel")

        layout = QVBoxLayout(panel)
        layout.setContentsMargins(14, 14, 14, 14)
        layout.setSpacing(10)

        title = QLabel("QUEUE NEW TASK")
        title.setObjectName("sectionTitle")

        selector_layout = QHBoxLayout()

        self.sender_selector = QComboBox()
        self.recipient_selector = QComboBox()
        self.task_type_selector = QComboBox()

        self.task_type_selector.addItems(
            [
                "message",
                "report",
                "email",
                "sms",
                "backup",
            ]
        )

        selector_layout.addWidget(
            QLabel("Sender:")
        )
        selector_layout.addWidget(
            self.sender_selector,
            1,
        )

        selector_layout.addWidget(
            QLabel("Recipient:")
        )
        selector_layout.addWidget(
            self.recipient_selector,
            1,
        )

        selector_layout.addWidget(
            QLabel("Type:")
        )
        selector_layout.addWidget(
            self.task_type_selector,
            1,
        )

        message_layout = QHBoxLayout()

        self.task_input = QTextEdit()
        self.task_input.setPlaceholderText(
            "Enter a message or task..."
        )
        self.task_input.setMaximumHeight(80)

        self.queue_task_button = QPushButton(
            "QUEUE TASK"
        )
        self.queue_task_button.setObjectName(
            "primaryButton"
        )
        self.queue_task_button.setMinimumWidth(140)

        message_layout.addWidget(
            self.task_input,
            1,
        )
        message_layout.addWidget(
            self.queue_task_button,
        )

        layout.addWidget(title)
        layout.addLayout(selector_layout)
        layout.addLayout(message_layout)

        return panel

    # ------------------------------------------------------------
    # Signals
    # ------------------------------------------------------------

    def connect_signals(self):
        self.refresh_button.clicked.connect(
            self.refresh_all
        )

        self.add_member_button.clicked.connect(
            self.open_add_member_dialog
        )

        self.queue_task_button.clicked.connect(
            self.queue_task
        )

    # ------------------------------------------------------------
    # Refreshing data
    # ------------------------------------------------------------

    def refresh_all(self):
        members_loaded = self.refresh_members()

        if members_loaded:
            self.refresh_tasks()

    def refresh_members(self):
        try:
            self.members = self.api.get_members()
        except APIError as api_error:
            self.set_server_status(False)
            self.show_status(str(api_error))
            return False

        self.set_server_status(True)

        selected_sender = (
            self.sender_selector.currentData()
        )
        selected_recipient = (
            self.recipient_selector.currentData()
        )

        self.member_list.clear()
        self.sender_selector.clear()
        self.recipient_selector.clear()

        for member in self.members:
            member_id = member["id"]
            member_name = member["name"]
            role = member.get("role", "Member")

            list_item = QListWidgetItem(
                f"{member_name}\n{role}"
            )
            list_item.setData(
                Qt.UserRole,
                member_id,
            )

            self.member_list.addItem(list_item)

            self.sender_selector.addItem(
                member_name,
                member_id,
            )

            self.recipient_selector.addItem(
                member_name,
                member_id,
            )

        self.restore_selector_value(
            self.sender_selector,
            selected_sender,
        )

        self.restore_selector_value(
            self.recipient_selector,
            selected_recipient,
        )

        return True

    def refresh_tasks(self):
        try:
            tasks = self.api.get_tasks()
        except APIError as api_error:
            self.set_server_status(False)
            self.show_status(str(api_error))
            return

        self.set_server_status(True)
        self.task_list.clear()

        for task in reversed(tasks):
            item = QListWidgetItem(
                self.format_task(task)
            )

            self.task_list.addItem(item)

        self.task_count_label.setText(
            f"{len(tasks)} TASKS"
        )

    def restore_selector_value(
        self,
        selector,
        member_id,
    ):
        if member_id is None:
            return

        index = selector.findData(member_id)

        if index >= 0:
            selector.setCurrentIndex(index)

    # ------------------------------------------------------------
    # Member actions
    # ------------------------------------------------------------

    def open_add_member_dialog(self):
        dialog = AddMemberDialog(self)
        dialog.setStyleSheet(CHIMERA_STYLE)

        if dialog.exec() != QDialog.Accepted:
            return

        data = dialog.member_data()

        if not data["name"]:
            self.show_status(
                "Member name is required."
            )
            return

        try:
            created_member = self.api.create_member(
                name=data["name"],
                contact=data["contact"],
                role=data["role"],
            )
        except APIError as api_error:
            self.show_status(str(api_error))
            return

        self.refresh_members()

        self.show_status(
            f"Added member: "
            f"{created_member['name']}"
        )

    # ------------------------------------------------------------
    # Task actions
    # ------------------------------------------------------------

    def queue_task(self):
        sender_id = self.sender_selector.currentData()
        recipient_id = (
            self.recipient_selector.currentData()
        )

        payload = self.task_input.toPlainText().strip()
        task_type = (
            self.task_type_selector.currentText()
        )

        if sender_id is None or recipient_id is None:
            self.show_status(
                "Add at least two members first."
            )
            return

        if sender_id == recipient_id:
            self.show_status(
                "Sender and recipient must differ."
            )
            return

        if not payload:
            self.show_status(
                "Enter a task or message."
            )
            self.task_input.setFocus()
            return

        try:
            created_task = self.api.create_task(
                sender_id=sender_id,
                recipient_id=recipient_id,
                task_type=task_type,
                payload=payload,
            )
        except APIError as api_error:
            self.show_status(str(api_error))
            return

        self.task_input.clear()
        self.refresh_tasks()

        self.show_status(
            f"Queued task {created_task['id']}."
        )

    # ------------------------------------------------------------
    # Display helpers
    # ------------------------------------------------------------

    def format_task(self, task):
        sender = task.get(
            "sender_name",
            "Unknown Member",
        )

        recipient = task.get(
            "recipient_name",
            "Unknown Member",
        )

        payload = task.get("payload", "")
        task_type = task.get("type", "message")
        status = task.get("status", "unknown")
        task_id = task.get("id", "?")
        attempts = task.get("attempts", 0)

        return (
            f"#{task_id}  {sender} → {recipient}\n"
            f"{payload}\n"
            f"Type: {task_type}   "
            f"Status: {status}   "
            f"Attempts: {attempts}"
        )

    def set_server_status(self, online):
        if online:
            self.server_status.setText(
                "● SERVER ONLINE"
            )
            object_name = "serverOnline"
        else:
            self.server_status.setText(
                "● SERVER OFFLINE"
            )
            object_name = "serverOffline"

        self.server_status.setObjectName(
            object_name
        )

        self.server_status.style().unpolish(
            self.server_status
        )
        self.server_status.style().polish(
            self.server_status
        )

    def show_status(self, message):
        self.statusBar().showMessage(
            message,
            6000,
        )
        
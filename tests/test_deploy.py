import importlib.util
import tempfile
import unittest
from contextlib import redirect_stdout
from io import StringIO
from pathlib import Path
from unittest.mock import patch


spec = importlib.util.spec_from_file_location(
    "deploy", Path(__file__).resolve().parents[1] / "deploy.py"
)
deploy = importlib.util.module_from_spec(spec)
spec.loader.exec_module(deploy)


class SkillDeploymentTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.files = self.root / "files"
        self.source = self.files / "claude" / "skills" / "example"
        self.source.mkdir(parents=True)
        (self.source / "SKILL.md").write_text(
            "---\nname: example\ndescription: Test skill.\n---\nShared instructions.\n"
        )
        self.claude = self.root / "home" / ".claude" / "skills"
        self.codex = self.root / "home" / ".agents" / "skills"
        for name, value in (
            ("FILES", self.files),
            ("SKILL_DIRS", (self.claude, self.codex)),
        ):
            patcher = patch.object(deploy, name, value)
            patcher.start()
            self.addCleanup(patcher.stop)
        output = redirect_stdout(StringIO())
        output.__enter__()
        self.addCleanup(output.__exit__, None, None, None)

    def test_both_agents_link_to_the_same_folder(self):
        deploy.cmd_skills()
        for root in (self.claude, self.codex):
            self.assertTrue((root / "example").is_symlink())
            self.assertEqual((root / "example").resolve(), self.source)

    def test_new_supporting_files_are_visible_without_redeploy(self):
        deploy.cmd_skills()
        sidecar = self.source / "agents" / "openai.yaml"
        sidecar.parent.mkdir()
        sidecar.write_text("policy:\n  allow_implicit_invocation: false\n")
        for root in (self.claude, self.codex):
            self.assertEqual((root / "example/agents/openai.yaml").read_text(), sidecar.read_text())

    def test_old_skill_folders_are_backed_up_outside_discovery(self):
        for root in (self.claude, self.codex):
            old = root / "example"
            old.mkdir(parents=True)
            (old / "SKILL.md").write_text("Old instructions.")
            existing = root.parent / "skill-backups" / "example.bak"
            existing.mkdir(parents=True)
            (existing / "SKILL.md").write_text("Earlier backup.")
        deploy.cmd_skills()
        for root in (self.claude, self.codex):
            backups = root.parent / "skill-backups"
            self.assertEqual((backups / "example.bak/SKILL.md").read_text(), "Earlier backup.")
            self.assertEqual((backups / "example.bak.1/SKILL.md").read_text(), "Old instructions.")
            self.assertEqual([p.name for p in root.iterdir()], ["example"])

    def test_repeat_deployment_does_not_replace_links_or_create_backups(self):
        deploy.cmd_skills()
        inodes = [(root / "example").lstat().st_ino for root in (self.claude, self.codex)]
        deploy.cmd_skills()
        self.assertEqual(inodes, [(root / "example").lstat().st_ino for root in (self.claude, self.codex)])
        for root in (self.claude, self.codex):
            self.assertFalse((root.parent / "skill-backups").exists())

    def test_dry_run_does_not_create_directories(self):
        deploy.cmd_skills(dry_run=True)
        self.assertFalse((self.root / "home").exists())

    def test_dry_run_does_not_change_existing_skill(self):
        old = self.codex / "example"
        old.mkdir(parents=True)
        (old / "SKILL.md").write_text("Old instructions.")
        deploy.cmd_skills(dry_run=True)
        self.assertFalse(old.is_symlink())
        self.assertEqual((old / "SKILL.md").read_text(), "Old instructions.")
        self.assertFalse((self.codex.parent / "skill-backups").exists())

    def test_hidden_and_non_skill_folders_are_ignored(self):
        base = self.source.parent
        (base / "notes").mkdir()
        (base / "notes/reference.md").write_text("Not a skill.")
        (base / ".hidden").mkdir()
        (base / ".hidden/SKILL.md").write_text("Hidden skill.")
        (base / "README.md").write_text("Skill documentation.")
        deploy.cmd_skills()
        for root in (self.claude, self.codex):
            self.assertEqual([p.name for p in root.iterdir()], ["example"])

    def test_unmanaged_and_system_skills_are_preserved(self):
        unmanaged = self.codex / "other"
        unmanaged.mkdir(parents=True)
        (unmanaged / "SKILL.md").write_text("Unmanaged skill.")
        system = self.codex.parent.parent / ".codex" / "skills" / ".system"
        system.mkdir(parents=True)
        (system / "sentinel").write_text("Built-in skill.")
        deploy.cmd_skills()
        self.assertEqual((unmanaged / "SKILL.md").read_text(), "Unmanaged skill.")
        self.assertEqual((system / "sentinel").read_text(), "Built-in skill.")

    def test_link_command_deploys_skills_separately_from_config_files(self):
        (self.files / "claude/settings.json").write_text("{}\n")
        (self.files / "codex").mkdir()
        (self.files / "codex/config.toml").write_text('model = "test"\n')
        codex_home = self.codex.parent.parent / ".codex"
        with patch.object(deploy, "FILE_LINKS", {}), patch.object(
            deploy, "DIR_LINKS", {"claude": self.claude.parent, "codex": codex_home}
        ):
            deploy.cmd_link()
        self.assertTrue((self.claude.parent / "settings.json").is_symlink())
        self.assertTrue((codex_home / "config.toml").is_symlink())
        self.assertTrue((self.claude / "example").is_symlink())
        self.assertTrue((self.codex / "example").is_symlink())
        self.assertFalse((codex_home / "skills").exists())

    def test_broken_backup_symlinks_are_not_overwritten(self):
        destination = self.root / "target"
        destination.write_text("Old file.")
        backup = self.root / "target.bak"
        backup.symlink_to(self.root / "missing")
        deploy.link_file(self.source / "SKILL.md", destination, False)
        self.assertTrue(backup.is_symlink())
        self.assertEqual((self.root / "target.bak.1").read_text(), "Old file.")

    def test_broken_destination_symlinks_can_be_replaced(self):
        self.codex.mkdir(parents=True)
        (self.codex / "example").symlink_to(self.root / "missing")
        deploy.cmd_skills()
        self.assertEqual((self.codex / "example").resolve(), self.source)


if __name__ == "__main__":
    unittest.main()

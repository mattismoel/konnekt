BEGIN TRANSACTION;

-- INSERT TEAMS --
INSERT OR IGNORE INTO team (id, name, display_name, description) VALUES
(1, 'admin', 'Administrator', 'Administrator af hjemmesiden.'),
(2, 'member', 'Medlem', 'Medlem af foreningen.'),
(3, 'project-leader', 'Projektleder', 'Projektleder af foreningen.'),
(4, 'booking', 'Booking', 'Booking af kunstnere.'),
(5, 'public-relations', 'PR', 'Håndtering af foreningens offentlige og medie-mæssige tilstedeværelse.'),
(6, 'visual-identity', 'Visuel Identitet', 'Håndtering af foreningens visuelle identitet og design.'),
(7, 'event-management', 'Event-management', 'Håndtering, planlægning og afvikling foreningens events'),
(8, 'economy', 'Økonomi', 'Håndtering af foreningens økonomi.');

-- INSERT PERMISSIONS --
INSERT OR IGNORE INTO permission (id, name, display_name, description) VALUES
(1, 'view:event', 'Events (se)', 'Tilladeslse til at se events.'),
(2, 'edit:event', 'Events (redigér)', 'Tilladelse til at redigére events.'),
(3, 'delete:event', 'Events (slet)', 'Tilladelse til at slette events.'),

(4, 'view:artist', 'Kunstnere (se)', 'Tilladeslse til at se kunstnere.'),
(5, 'edit:artist', 'Kunstnere (redigér)', 'Tilladelse til at redigére kunstnere.'),
(6, 'delete:artist', 'Kunstnere (slet)', 'Tilladelse til at slette kunstnere.'),

(7, 'view:venue', 'Venues (se)', 'Tilladeslse til at se venues.'),
(8, 'edit:venue', 'Venues (redigér)', 'Tilladelse til at redigére venues.'),
(9, 'delete:venue', 'Venues (slet)', 'Tilladelse til at slette venues.'),

(10, 'view:genre', 'Genrer (se)', 'Tilladeslse til at se genrer.'),
(11, 'edit:genre', 'Genrer (redigér)', 'Tilladelse til at redigére genrer.'),
(12, 'delete:genre', 'Genrer (slet)', 'Tilladelse til at slette genrer.'),

(13, 'view:content', 'Indhold (se)', 'Tilladeslse til at se hjemmesideindhold.'),
(14, 'edit:content', 'Indhold (redigér)', 'Tilladelse til at redigére hjemmesideindhold.'),
(15, 'delete:content', 'Indhold (slet)', 'Tilladelse til at slette hjemmesideindhold.'),

(16, 'view:member', 'Medlemmer (se)', 'Tilladeslse til at se medlemmer.'),
(17, 'edit:member', 'Medlemmer (redigér)', 'Tilladelse til at redigére medlemmer.'),
(18, 'delete:member', 'Medlemmer (slet)', 'Tilladelse til at slette medlemmer.'),

(22, 'view:team', 'Hold (se)', 'Tilladeslse til at se hold.'),
(23, 'edit:team', 'Hold (redigér)', 'Tilladelse til at redigére hold.'),
(24, 'delete:team', 'Hold (slet)', 'Tilladelse til at slette hold.'),

(25, 'view:permission', 'Tilladelser (se)', 'Tilladeslse til at se tilladelse.'),
(26, 'edit:permission', 'Tilladelser (redigér)', 'Tilladelse til at redigére tilladelser.'),
(27, 'delete:permission', 'Tilladelser (slet)', 'Tilladelse til at slette tilladelser.');

-- ADMIN PERMISSIONS (All permissions).
INSERT OR IGNORE INTO teams_permissions (team_id, permission_id) SELECT 1, id FROM permission;

-- MEMBER PERMISSIONS (view-only).
INSERT OR IGNORE INTO teams_permissions (team_id, permission_id) SELECT 2, id FROM permission
WHERE name IN (
  'view:event',
  'view:artist',
  'view:venue',
  'view:genre',
  'view:content',
  'view:member',
  'view:team',
  'view:permission'
);

-- PROJECT LEADER PERMISSIONS.
INSERT OR IGNORE INTO teams_permissions (team_id, permission_id) SELECT 3, id FROM permission
WHERE name IN (
  'view:event', 'edit:event', 'delete:event',
  'view:artist', 'edit:artist', 'delete:artist'
);

-- BOOKING PERMISSIONS.
INSERT OR IGNORE INTO teams_permissions (team_id, permission_id) SELECT 4, id FROM permission
WHERE name IN (
  'view:artist', 'edit:artist', 'delete:artist',
  'view:venue', 'edit:venue', 'delete:venue',
  'view:genre', 'edit:genre', 'delete:genre'
);


-- PUBLIC RELATIONS PERMISSIONS.
INSERT OR IGNORE INTO teams_permissions (team_id, permission_id) SELECT 5, id FROM permission
WHERE name IN (
  'view:event', 'edit:event', 'delete:event',
  'view:venue', 'edit:venue', 'delete:venue',
  'view:content', 'edit:content', 'delete:content'
);

-- VISUAL IDENTITY PERMISSIONS.
INSERT OR IGNORE INTO teams_permissions (team_id, permission_id) SELECT 6, id FROM permission
WHERE name IN (
  'view:event', 'edit:event',
  'view:artist', 'edit:artist',
  'view:content', 'edit:content', 'delete:content'
);

COMMIT;

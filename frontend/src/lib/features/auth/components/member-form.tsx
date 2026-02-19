import {
  Controller,
  FormProvider,
  useForm,
  useFormContext,
} from "react-hook-form";

import { zodResolver } from "@hookform/resolvers/zod";

import {
  deleteMember,
  editMember,
  memberForm,
  type Member,
  type MemberFormValues,
} from "../member";

import type { Team } from "../team";
import { useAuth } from "@/lib/context/auth";
import ProfilePictureSelector from "@/lib/components/profile-picture-selector";
import FormField from "@/lib/components/form-field";
import Input from "@/lib/components/ui/input";
import Button from "@/lib/components/ui/button/button";
import { APIError, createSubmitHandler } from "@/lib/api";
import { createContext, useContext, useState } from "react";
import PillList from "@/lib/components/pill-list";
import { FaPen } from "react-icons/fa6";
import { useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "@tanstack/react-router";
import Modal from "@/lib/components/ui/modal";
import type { Entry as EntryType } from "@/lib/components/ui/picker/entry";
import MultiPicker from "@/lib/components/ui/picker/multi-picker";
import Searchbar from "@/lib/components/searchbar";
import { formatRelative } from "date-fns";
import { useSearch } from "@/lib/hooks/useSearch";
import { useToast } from "@/lib/context/toast";

type MemberFormContext = {
  member: Member;
  memberTeams: Team[];
  isCurrentMember: boolean;

  isEditable: boolean;

  teams: Team[];
};

const MemberFormContext = createContext<MemberFormContext | undefined>(
  undefined,
);

const useMemberFormContext = () => {
  const ctx = useContext(MemberFormContext);
  if (!ctx) throw new Error("No MemberFormContext.Provider found!");

  return ctx;
};

type Props = {
  member: Member;
  memberTeams: Team[];

  teams: Team[];
};

const MemberForm = ({ member, memberTeams, teams }: Props) => {
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const { hasPermissions, member: currentMember } = useAuth();

  const isEditable = hasPermissions(["edit:member"]);
  const isCurrentMember = currentMember?.id === member.id;

  const methods = useForm({
    defaultValues: {
      image: undefined,
      firstName: member.firstName,
      lastName: member.lastName,
      email: member.email,
      specialRole: member.specialRole,
      memberTeams: memberTeams.map(({ id }) => id),
    },
    resolver: zodResolver(memberForm),
  });

  const {
    formState: { errors, isDirty },
    control,
    handleSubmit,
  } = methods;

  let fullName = `${member.firstName} ${member.lastName}`;

  const onSubmit = createSubmitHandler({
    action: async (form: MemberFormValues) => {
      await editMember(member.id, form);
      await queryClient.invalidateQueries({ queryKey: ["members"] });
      navigate({ to: "/admin/members" });
    },
    errorMessage: "Kunne ikke redigere medlem",
    successMessage: "Medlem redigeret",
    navigateTo: "/admin/members",
  });

  const { addToast } = useToast();

  const handleDeleteMember = async () => {
    if (
      !confirm(
        `Er du sikker på at du vil slette ${fullName} fra foreningen?\n\nHandlingen kan ikke fortrydes.`,
      )
    ) {
      return;
    }
    try {
      await deleteMember(member.id);
      addToast("Medlem slettet");
      await queryClient.invalidateQueries({ queryKey: ["members"] });
    } catch (e) {
      if (e instanceof APIError) {
        addToast("Kunne ikke slette medlemmet", e.cause, "error");
        throw e;
      }
      addToast("Kunne ikke slette medlemmet", "Noget gik galt...", "error");
      throw e;
    }
  };

  return (
    <MemberFormContext.Provider
      value={{ member, teams, memberTeams, isCurrentMember, isEditable }}
    >
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-8">
          <div className="flex w-full flex-col items-center gap-8">
            <Controller
              control={control}
              name="image"
              render={({ field }) => (
                <FormField error={errors.image} className="justify-center">
                  <ProfilePictureSelector
                    {...field}
                    src={member.profilePictureUrl}
                  />
                </FormField>
              )}
            />
            <div className="flex flex-col items-center space-y-4 md:items-start">
              <div className="flex flex-col items-center space-y-1">
                <h1 className="text-2xl font-semibold text-heading">
                  {fullName}
                </h1>
                <span className="text-text/75 text-center md:text-left">
                  {memberTeams.map(({ displayName }) => displayName).join(", ")}
                </span>
              </div>
            </div>
          </div>

          <GeneralSection />
          <TeamsSection />

          {(isEditable || isCurrentMember) && (
            <div className="flex flex-col gap-2">
              <Button type="submit" disabled={!isDirty} className="w-full">
                Opdatér
              </Button>
              <Button
                variant="dangerous"
                className="w-full"
                onClick={handleDeleteMember}
              >
                Slet
              </Button>
            </div>
          )}
          <span className="text-text/50">
            Medlem siden {formatRelative(member.createdAt, new Date())}
          </span>
        </form>
      </FormProvider>
    </MemberFormContext.Provider>
  );
};

const GeneralSection = () => {
  const {
    formState: { errors },
    register,
  } = useFormContext<MemberFormValues>();
  const { isCurrentMember } = useMemberFormContext();

  return (
    <section>
      <h1 className="mb-4 font-heading text-2xl font-bold text-heading">
        Generelt
      </h1>

      <div className="flex flex-col gap-4">
        <div className="flex gap-4">
          <FormField error={errors.firstName}>
            <Input
              {...register("firstName")}
              placeholder="Fornavn"
              disabled={!isCurrentMember}
            />
          </FormField>
          <FormField error={errors.lastName}>
            <Input
              {...register("lastName")}
              placeholder="Efternavn"
              disabled={!isCurrentMember}
            />
          </FormField>
        </div>

        <FormField error={errors.email}>
          <Input
            {...register("email")}
            type="email"
            placeholder="Email"
            disabled={!isCurrentMember}
          />
        </FormField>
      </div>
    </section>
  );
};

const TeamsSection = () => {
  const {
    control,
    formState: { errors },
    register,
  } = useFormContext<MemberFormValues>();
  const { teams, isEditable } = useMemberFormContext();

  const [showPicker, setShowPicker] = useState(false);

  const entries: EntryType[] = teams.map(({ id, displayName }) => ({
    id: id.toString(),
    value: id.toString(),
    name: displayName,
  }));

  const { search, results, setSearch } = useSearch(entries, "name");

  return (
    <section>
      <h1 className="mb-8 font-heading text-2xl font-bold text-heading">
        Hold
      </h1>

      <div className="flex flex-col gap-8">
        <FormField error={errors.specialRole}>
          <Input
            placeholder="Specialtitel"
            {...register("specialRole")}
            disabled={!isEditable}
          />
        </FormField>

        <Controller
          control={control}
          name="memberTeams"
          render={({ field: { value, onChange }, fieldState: { error } }) => {
            const selectedEntries = entries.filter((e) =>
              value?.includes(parseInt(e.value)),
            );

            return (
              <>
                <PillList entries={selectedEntries.map((e) => e.name)}>
                  {isEditable && (
                    <Button
                      variant="primary"
                      onClick={() => setShowPicker(true)}
                      className="h-10"
                    >
                      <FaPen />
                      Vælg
                    </Button>
                  )}
                </PillList>

                <FormField error={error}>
                  <Modal show={showPicker} onClose={() => setShowPicker(false)}>
                    <Modal.Header>
                      <Modal.Title>Vælg medlemshold...</Modal.Title>
                      <Modal.Description>
                        Her kan du vælge de medlemshold, som medlemmet
                        associeres med.
                      </Modal.Description>
                    </Modal.Header>
                    <Modal.Content className="flex flex-col gap-4">
                      <Searchbar
                        search={search}
                        onChange={(s) => setSearch(s)}
                      />
                      <MultiPicker
                        selected={selectedEntries}
                        entries={results}
                        onChange={(newEntries) =>
                          onChange(newEntries.map((e) => parseInt(e.value)))
                        }
                      />
                    </Modal.Content>
                    <Modal.Footer>
                      <Button
                        type="button"
                        onClick={() => setShowPicker(false)}
                      >
                        Vælg
                      </Button>
                    </Modal.Footer>
                  </Modal>
                </FormField>
              </>
            );
          }}
        />
      </div>
    </section>
  );
};

export default MemberForm;
